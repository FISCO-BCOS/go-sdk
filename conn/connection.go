// Copyright FISCO-BCOS go-sdk
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conn

import "C"
import (
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/FISCO-BCOS/bcos-c-sdk/bindings/go/csdk"
	"github.com/FISCO-BCOS/crypto/tls"
	"github.com/FISCO-BCOS/crypto/x509"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

var (
	ErrClientQuit                = errors.New("client is closed")
	ErrNoResult                  = errors.New("no result in JSON-RPC response")
	ErrNoRpcMehtod               = errors.New("no rpc method")
	ErrSubscriptionQueueOverflow = errors.New("subscription queue overflow")
	errClientReconnected         = errors.New("client reconnected")
	errDead                      = errors.New("connection lost")
)

const (
	// Timeouts
	tcpKeepAliveInterval = 30 * time.Second
	defaultDialTimeout   = 10 * time.Second // used if context has no deadline
	subscribeTimeout     = 5 * time.Second  // overall timeout eth_subscribe, rpc_modules calls
	amopTimeout          = 1000
)

const (
	// Subscriptions are removed when the subscriber cannot keep up.
	//
	// This can be worked around by supplying a channel with sufficiently sized buffer,
	// but this can be inconvenient and hard to explain in the docs. Another issue with
	// buffered channels is that the buffer is static even though it might not be needed
	// most of the time.
	//
	// The approach taken here is to maintain a per-subscription linked list buffer
	// shrinks on demand. If the buffer reaches the size below, the subscription is
	// dropped.
	maxClientSubscriptionBuffer = 20000
)

// Error wraps RPC errors, which contain an error code in addition to the message.
type Error interface {
	Error() string  // returns the message
	ErrorCode() int // returns the code
}

// ServerCodec implements reading, parsing and writing RPC messages for the server side of
// a RPC session. Implementations must be go-routine safe since the codec can be called in
// multiple go-routines concurrently.
type ServerCodec interface {
	Read() (msgs []*jsonrpcMessage, isBatch bool, err error)
	Close()
	jsonWriter
}

// jsonWriter can write JSON messages to its underlying connection.
// Implementations must be safe for concurrent use.
type jsonWriter interface {
	Write(context.Context, interface{}) error
	// Closed returns a channel which is closed when the connection is closed.
	Closed() <-chan interface{}
	// RemoteAddr returns the peer address of the connection.
	RemoteAddr() string
}

// Connection represents a connection to an RPC server.
type Connection struct {
	csdk      *csdk.CSDK
	idCounter uint32

	// This function, if non-nil, is called when the connection is lost.
	reconnectFunc reconnectFunc

	// writeConn is used for writing to the connection on the caller's goroutine. It should
	// only be accessed outside of dispatch, with the write lock held. The write lock is
	// taken by sending on requestOp and released by sending on sendDone.
	writeConn jsonWriter

	// for dispatch
	close       chan struct{}
	closing     chan struct{}    // closed when client is quitting
	didClose    chan struct{}    // closed when client quits
	reconnected chan ServerCodec // where write/reconnect sends the new connection
	readOp      chan readOp      // read messages
	readErr     chan error       // errors from read
	reqInit     chan *requestOp  // register response IDs, takes write lock
	reqSent     chan error       // signals write completion, releases write lock
	reqTimeout  chan *requestOp  // removes response IDs when call timeout expires
}

type reconnectFunc func(ctx context.Context) (ServerCodec, error)

type clientContextKey struct{}

type clientConn struct {
	codec ServerCodec
}

func (c *Connection) newClientConn(conn ServerCodec) *clientConn {
	return &clientConn{conn}
}

func (cc *clientConn) close(err error, inflightReq *requestOp) {
	cc.codec.Close()
}

type readOp struct {
	msgs  []*jsonrpcMessage
	batch bool
}

type requestOp struct {
	ids  []json.RawMessage
	err  error
	resp chan *jsonrpcMessage // receives up to len(ids) responses
	//respData     chan *resRpcMessage  // receives up to len(ids) responses
	respChanData *csdk.CallbackChan
}

type EventLogRespResult struct {
	LogIndex         int    `json:"logIndex"`
	TransactionIndex int    `json:"transactionIndex"`
	TransactionHash  string `json:"transactionHash"`
	//BlockHash        string   `json:"blockHash"`
	BlockNumber uint64   `json:"blockNumber"`
	Address     string   `json:"address"`
	Data        string   `json:"data"`
	Topics      []string `json:"topics"`
}

type eventLogResp struct {
	FilterID string               `json:"id"`
	Result   []EventLogRespResult `json:"result"`
	Status   int                  `json:"status"`
}

func (op *requestOp) waitRpcMessage(ctx context.Context) (*jsonrpcMessage, interface{}, error) {
	respBody := <-op.respChanData.Data
	var respData jsonrpcMessage
	if err := json.Unmarshal(respBody.Result, &respData); err != nil {
		//log.Println("json unmarshal res body:", respBody)
		//log.Println("json unmarshal err:", err)
		return nil, nil, err
	}
	return &respData, respData.Result, op.err
}

func processEventLogMsg(respBody []byte, handler interface{}) {
	var eventLogResponse eventLogResp
	err := json.Unmarshal(respBody, &eventLogResponse)
	if err != nil {
		logrus.Warnf("unmarshal eventLogResponse failed, err: %v\n", err)
		return
	}
	if len(eventLogResponse.Result) == 0 {
		return
	}
	logs := []types.Log{}
	for _, eventLog := range eventLogResponse.Result {
		number := eventLog.BlockNumber
		logIndex := eventLog.LogIndex
		txIndex := eventLog.TransactionIndex
		topics := []common.Hash{}
		for _, topic := range eventLog.Topics {
			topics = append(topics, common.HexToHash(topic))
		}
		data := common.FromHex(eventLog.Data)
		logs = append(logs, types.Log{
			Address:     common.HexToAddress(eventLog.Address),
			Topics:      topics,
			Data:        data,
			BlockNumber: uint64(number),
			TxHash:      common.HexToHash(eventLog.TransactionHash),
			TxIndex:     uint(txIndex),
			//BlockHash:   common.HexToHash(eventLog.BlockHash),
			Index:   uint(logIndex),
			Removed: false,
		})
	}
	eventHander := handler.(func(int, []types.Log))
	go eventHander(eventLogResponse.Status, logs)
}

func (op *requestOp) waitMessage(ctx context.Context, c *Connection, method string, handler interface{}) error {
	for {
		select {
		case <-ctx.Done():
			//switch method {
			//	case "subscribeEventLogs":
			//		taskId := ctx.Value("taskId").(string)
			//		c.csdk.UnsubscribeEvent(op.respChanData, taskId)
			//}
			return ctx.Err()
		case respBody := <-op.respChanData.Data:
			switch method {
			case "subscribeEventLogs":
				processEventLogMsg(respBody.Result, handler)
			case "subscribeTopic":
				handler.(func([]byte, *[]byte))(respBody.Result, nil)
			case "SendAMOPMsg":
				handler.(func([]byte, *[]byte))(respBody.Result, nil)
			case "broadcastAMOPMsg":
				handler.(func([]byte, *[]byte))(respBody.Result, nil)
			case "subscribeBlockNumberNotify":
				//log.Println("subscribeBlockNumberNotify respBody:",respBody)
				if respBody.Err == nil {
					blockNum := int64(binary.LittleEndian.Uint64(respBody.Result))
					handler.(func(int64))(blockNum)
				}
			default:
				return ErrNoResult
			}
		}
	}
	return nil
}

// DialContextChannel creates a new Channel client, just like Dial.
func DialContextChannel(rawurl string, caRoot, certContext, keyContext []byte, groupID int) (*Connection, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caRoot)
	if !ok {
		return nil, errors.New("failed to parse root certificate")
	}
	cer, err := tls.X509KeyPair(certContext, keyContext)
	if err != nil {
		return nil, err
	}
	config := &tls.Config{RootCAs: roots, Certificates: []tls.Certificate{cer}, MinVersion: tls.VersionTLS12, PreferServerCipherSuites: true,
		InsecureSkipVerify: true}
	config.CurvePreferences = append(config.CurvePreferences, tls.CurveSecp256k1, tls.CurveP256)
	return DialChannelWithClient(rawurl, config, groupID)
}

// ClientFromContext Connection retrieves the client from the context, if any. This can be used to perform
// 'reverse calls' in a handler method.
func ClientFromContext(ctx context.Context) (*Connection, bool) {
	client, ok := ctx.Value(clientContextKey{}).(*Connection)
	return client, ok
}

func NewClient(connect reconnectFunc, csdk *csdk.CSDK) (*Connection, error) {
	c := initClient(nil)
	c.reconnectFunc = connect
	c.csdk = csdk
	return c, nil
}

func newClient(initctx context.Context, connect reconnectFunc) (*Connection, error) {
	conn, err := connect(initctx)
	if err != nil {
		return nil, err
	}
	c := initClient(conn)
	c.reconnectFunc = connect
	return c, nil
}

func initClient(conn ServerCodec) *Connection {
	c := &Connection{
		writeConn:   conn,
		close:       make(chan struct{}),
		closing:     make(chan struct{}),
		didClose:    make(chan struct{}),
		reconnected: make(chan ServerCodec),
		readOp:      make(chan readOp),
		readErr:     make(chan error),
		reqInit:     make(chan *requestOp),
		reqSent:     make(chan error, 1),
		reqTimeout:  make(chan *requestOp),
	}
	return c
}

func (c *Connection) nextID() json.RawMessage {
	id := atomic.AddUint32(&c.idCounter, 1)
	return strconv.AppendUint(nil, uint64(id), 10)
}

// Close closes the client, aborting any in-flight requests.
func (c *Connection) Close() {
	c.csdk.Close()
}

// Close closes the client, aborting any in-flight requests.
func (c *Connection) ReConn() {
	hc := c.writeConn.(*channelSession)
	hc.Reconnection()
}

// Call performs a JSON-RPC call with the given arguments and unmarshals into
// result if no error occurred.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
func (c *Connection) Call(result interface{}, method string, args ...interface{}) error {
	ctx := context.Background()
	return c.CallContext(ctx, result, method, args...)
}

// CallContext performs a JSON-RPC call with the given arguments. If the context is
// canceled before the call has successfully returned, CallContext returns immediately.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
func (c *Connection) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	//logrus.Infof("CallContext method:%s\n", method)
	op := &requestOp{respChanData: &csdk.CallbackChan{Data: make(chan csdk.Response, 100)}}
	switch method {
	case "call":
		arg := args[1].(map[string]interface{})
		data := arg["data"].(string)
		to := arg["to"].(string)
		c.csdk.Call(op.respChanData, to, data)
	case "getGroupPeers":
		c.csdk.GetGroupPeers(op.respChanData)
	case "getPeers":
		c.csdk.GetPeers(op.respChanData)
	case "getBlockNumber":
		c.csdk.GetBlockNumber(op.respChanData)
	case "getBlockByNumber":
		blockNumber := args[1].(int64)
		onlyHeader := args[2].(bool)
		onlyTxHash := args[3].(bool)
		c.csdk.GetBlockByNumber(op.respChanData, blockNumber, onlyHeader, onlyTxHash)
	case "getBlockByHash":
		blockHash := args[1].(string)
		onlyHeader := args[2].(bool)
		onlyTxHash := args[3].(bool)
		c.csdk.GetBlockByHash(op.respChanData, blockHash, onlyHeader, onlyTxHash)
	case "getBlockHashByNumber":
		blockNumber := args[1].(int64)
		c.csdk.GetBlockHashByNumber(op.respChanData, blockNumber)
	case "getPbftView":
		c.csdk.GetPbftView(op.respChanData)
	case "getCode":
		address := args[1].(string)
		c.csdk.GetCode(op.respChanData, address)
	case "getSyncStatus":
		c.csdk.GetSyncStatus(op.respChanData)
	case "getConsensusStatus":
		c.csdk.GetConsensusStatus(op.respChanData)
	case "getSealerList":
		c.csdk.GetSealerList(op.respChanData)
	case "getObserverList":
		c.csdk.GetObserverList(op.respChanData)
	case "getTransactionReceipt":
		txHash := args[1].(string)
		withProof := args[2].(bool)
		c.csdk.GetTransactionReceipt(op.respChanData, txHash, withProof)
	case "getTransactionByHash":
		txHash := args[1].(string)
		withProof := args[2].(bool)
		c.csdk.GetTransaction(op.respChanData, txHash, withProof)
	case "getSystemConfigByKey":
		key := args[1].(string)
		c.csdk.GetSystemConfigByKey(op.respChanData, key)
	case "getTotalTransactionCount":
		c.csdk.GetTotalTransactionCount(op.respChanData)
	case "getGroupNodeInfo":
		c.csdk.GetGroupNodeInfo(op.respChanData)
	case "getGroupList":
		c.csdk.GetGroupList(op.respChanData)
	case "getGroupInfo":
		c.csdk.GetGroupInfo(op.respChanData)
	case "getGroupInfoList":
		c.csdk.GetGroupNodeInfoList(op.respChanData)
	case "getPendingTxSize":
		c.csdk.GetPendingTxSize(op.respChanData)
	case "sendRawTransaction":
		data := args[1].(string)
		contractAddress := args[2].(string)
		c.csdk.SendTransaction(op.respChanData, contractAddress, data, true)
	default:
		return ErrNoRpcMehtod
	}

	// dispatch has accepted the request and will close the channel when it quits.
	switch resp, _, err := op.waitRpcMessage(ctx); {
	case err != nil:
		return err
	case resp.Error != nil:
		return resp.Error
	case len(resp.Result) == 0:
		logrus.Errorf("result is null, %+v, err:%+v \n", resp, err)
		return ErrNoResult
	default:
		return json.Unmarshal(resp.Result, &result)
	}
}

func (c *Connection) CallHandlerContext(ctx context.Context, result interface{}, method string, topic string, reqData []byte, handler interface{}) error {
	//logrus.Infof("CallEventContext method:%s", method)
	op := &requestOp{respChanData: &csdk.CallbackChan{Data: make(chan csdk.Response, 100)}}
	switch method {
	case "subscribeTopic":
		c.csdk.SubscribeTopicWithCb(op.respChanData, topic)
	case "unsubscribeTopic":
		c.csdk.UnsubscribeTopicWithCb(op.respChanData, topic)
	case "SendAMOPMsg":
		c.csdk.PublishTopicMsg(op.respChanData, topic, reqData, amopTimeout)
	case "broadcastAMOPMsg":
		c.csdk.BroadcastAmopMsg(op.respChanData, topic, reqData)
	case "subscribeEventLogs":
		taskId := c.csdk.SubscribeEvent(op.respChanData, string(reqData))
		*result.(*string) = taskId
	case "unSubscribeEventLogs":
		c.csdk.UnsubscribeEvent(op.respChanData, string(reqData))
	case "subscribeBlockNumberNotify":
		c.csdk.RegisterBlockNotifier(op.respChanData)
	default:
		return ErrNoResult
	}

	var err error
	go func() {
		err = op.waitMessage(ctx, c, method, handler)
	}()
	return err
}

//func (c *Connection) AsyncSendTransaction(ctx context.Context, handler func(*types.Receipt, error), method string, args ...interface{}) error {
//	msg, err := c.newMessage(method, args...)
//	if err != nil {
//		return err
//	}
//	hc := c.writeConn.(*channelSession)
//	err = hc.asyncSendTransaction(msg, handler)
//	if err != nil {
//		return err
//	}
//	return nil
//}

//func (c *Connection) UnsubscribeBlockNumberNotify(groupID uint64) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.unSubscribeBlockNumberNotify(groupID)
//}

//func (c *Connection) newMessage(method string, paramsIn ...interface{}) (*jsonrpcMessage, error) {
//	msg := &jsonrpcMessage{Version: vsn, ID: c.nextID(), Method: method}
//	if paramsIn != nil { // prevent sending "params":null
//		var err error
//		if msg.Params, err = json.Marshal(paramsIn); err != nil {
//			return nil, err
//		}
//	}
//	return msg, nil
//}

func (c *Connection) reconnect(ctx context.Context) error {
	if c.reconnectFunc == nil {
		return errDead
	}

	if _, ok := ctx.Deadline(); !ok {
		var cancel func()
		ctx, cancel = context.WithTimeout(ctx, defaultDialTimeout)
		defer cancel()
	}
	newconn, err := c.reconnectFunc(ctx)
	if err != nil {
		// logrus.Trace("RPC client reconnect failed", "err", err)
		return err
	}
	select {
	case c.reconnected <- newconn:
		c.writeConn = newconn
		return nil
	case <-c.didClose:
		newconn.Close()
		return ErrClientQuit
	}
}

// drainRead drops read messages until an error occurs.
func (c *Connection) drainRead() {
	for {
		select {
		case <-c.readOp:
		case <-c.readErr:
			return
		}
	}
}

// GetBlockNumber returns BlockLimit
func (c *Connection) GetBlockNumber() int64 {
	hc, ok := c.writeConn.(*channelSession)
	if !ok {
		return 0
	}
	return hc.nodeInfo.blockNumber
}
