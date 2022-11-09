// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package conn

import "C"
import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
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

// BatchElem is an element in a batch request.
type BatchElem struct {
	Method string
	Args   []interface{}
	// The result is unmarshaled into this field. Result must be set to a
	// non-nil pointer value of the desired type, otherwise the response will be
	// discarded.
	Result interface{}
	// Error is set if the server returns an error for this request, or if
	// unmarshaling into Result fails. It is not set for I/O errors.
	Error error
}

// Connection represents a connection to an RPC server.
type Connection struct {
	isHTTP    bool
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
	respChanData *csdk.ChanData
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
	select {
	case respBody := <-op.respChanData.Data:
		var respData jsonrpcMessage
		if err := json.Unmarshal([]byte(respBody), &respData); err != nil {
			//log.Println("json unmarshal res body:", respBody)
			//log.Println("json unmarshal err:", err)
			return nil, nil, err
		}
		return &respData, respData.Result, op.err
	}
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
	for true {
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
				processEventLogMsg([]byte(respBody), handler)
			case "subscribeTopic":
				handler.(func([]byte, *[]byte))([]byte(respBody), nil)
			case "SendAMOPMsg":
				handler.(func([]byte, *[]byte))([]byte(respBody), nil)
			case "broadcastAMOPMsg":
				handler.(func([]byte, *[]byte))([]byte(respBody), nil)
			case "subscribeBlockNumberNotify":
				//log.Println("subscribeBlockNumberNotify respBody:",respBody)
				blockNum, err := strconv.Atoi(respBody)
				if err == nil {
					handler.(func(int64))(int64(blockNum))
				}
			default:
				return ErrNoResult
			}
		}
	}
	return nil
}

//func (op *requestOp) wait(ctx context.Context, c *Connection) (*resRpcMessage, interface{}, error) {
//	select {
//	// case <-ctx.Done():
//	// 	// Send the timeout to dispatch so it can remove the request IDs.
//	// 	// FIXME: remove the code below
//	// 	if !c.isHTTP {
//	// 		select {
//	// 		case c.reqTimeout <- op:
//	// 		case <-c.closing:
//	// 		}
//	// 	}
//	// 	return nil, ctx.Err()
//	case respBody := <-op.respChanData.Data:
//		var respData resRpcMessage
//		if err := json.Unmarshal([]byte(respBody), &respData); err != nil {
//			log.Println("json unmarshal res body:", respBody)
//			log.Println("json unmarshal err:", err)
//			return nil, nil, err
//		}
//		return &respData, respData.Result, op.err
//	}
//}

// DialContextHTTP creates a new RPC client, just like Dial.
//
// The context is used to cancel or time out the initial connection establishment. It does
// not affect subsequent interactions with the client.
func DialContextHTTP(rawurl string) (*Connection, error) {
	rawurl = strings.ToLower(rawurl)
	if !strings.Contains(rawurl, "http://") {
		rawurl = "http://" + rawurl
	}
	return DialHTTP(rawurl)
}

// DialContextChannel creates a new Channel client, just like Dial.
func DialContextChannel(rawurl string, caRoot, certContext, keyContext []byte, groupID int) (*Connection, error) {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM(caRoot)
	if !ok {
		//panic("failed to parse root certificate")
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
	_, isHTTP := conn.(*httpConn)
	c := &Connection{
		isHTTP:      isHTTP,
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
	//if c.isHTTP {
	//	return
	//}
	//hc := c.writeConn.(*channelSession)
	//hc.Close()
	c.csdk.Close()
}

// Close closes the client, aborting any in-flight requests.
func (c *Connection) ReConn() {
	if c.isHTTP {
		return
	}
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
	op := &requestOp{respChanData: &csdk.ChanData{Data: make(chan string, 100)}}
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
		c.csdk.GetBlockByNumber(op.respChanData, blockNumber, 0, 0)
	case "getBlockByHash":
		blockHash := args[1].(string)
		c.csdk.GetBlockByhash(op.respChanData, blockHash, 0, 0)
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
		c.csdk.GetTransactionReceipt(op.respChanData, txHash)
	case "getTransactionByHash":
		txHash := args[1].(string)
		c.csdk.GetTransaction(op.respChanData, txHash)
	case "getSystemConfigByKey":
		key := args[1].(string)
		c.csdk.GetSystemConfigByKey(op.respChanData, key)
	case "getTotalTransactionCount":
		c.csdk.GetTotalTransactionCount(op.respChanData)
	case "getGroupNodeInfo":
		nodeId := args[1].(string)
		c.csdk.GetGroupnodeInfo(op.respChanData, nodeId)
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
		c.csdk.SendTransaction(op.respChanData, contractAddress, data)
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

func (c *Connection) CallHandlerContext(ctx context.Context, result interface{}, method string, topic string, reqData string, handler interface{}) error {
	//logrus.Infof("CallEventContext method:%s", method)
	op := &requestOp{respChanData: &csdk.ChanData{Data: make(chan string, 100)}}
	switch method {
	case "subscribeTopic":
		c.csdk.SubscribeTopicWithCb(op.respChanData, topic)
	case "unsubscribeTopic":
		c.csdk.UnsubscribeTopicWithCb(op.respChanData, topic)
	case "SendAMOPMsg":
		c.csdk.PublishTopicMsg(op.respChanData, topic, reqData)
	case "broadcastAMOPMsg":
		c.csdk.BroadcastAmopMsg(op.respChanData, topic, reqData)
	case "subscribeEventLogs":
		taskId := c.csdk.SubscribeEvent(op.respChanData, reqData)
		*result.(*string) = taskId
	case "unSubscribeEventLogs":
		c.csdk.UnsubscribeEvent(op.respChanData, reqData)
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
//	if !c.isHTTP {
//		err = hc.asyncSendTransaction(msg, handler)
//	}
//	if err != nil {
//		return err
//	}
//	return nil
//}

// func (c *Connection) SubscribeEventLogs(eventLogParams types.EventLogParams, handler func(int, []types.Log)) (string, error) {
// 	hc := c.writeConn.(*channelSession)
// 	return hc.subscribeEvent(eventLogParams, handler)
// }

// func (c *Connection) UnSubscribeEventLogs(filterID string) error {
// 	hc := c.writeConn.(*channelSession)
// 	hc.eventLogMu.Lock()
// 	delete(hc.eventLogHandlers, filterID)
// 	hc.eventLogMu.Unlock()
// 	return nil
// }

// func (c *Connection) SubscribeTopic(topic string, handler func([]byte, *[]byte)) error {
// 	hc := c.writeConn.(*channelSession)
// 	return hc.subscribeTopic(topic, handler)
// }

//func (c *Connection) SubscribePrivateTopic(topic string, privateKey *ecdsa.PrivateKey, handler func([]byte, *[]byte)) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.subscribePrivateTopic(topic, privateKey, handler)
//}

//func (c *Connection) PublishPrivateTopic(topic string, publicKey []*ecdsa.PublicKey) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.publishPrivateTopic(topic, publicKey)
//}

//func (c *Connection) UnsubscribeTopic(topic string) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.unsubscribeTopic(topic)
//}

//func (c *Connection) UnsubscribePrivateTopic(topic string) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.unsubscribePrivateTopic(topic)
//}

//func (c *Connection) SendAMOPMsg(topic string, data []byte) ([]byte, error) {
//	hc := c.writeConn.(*channelSession)
//	return hc.sendAMOPMsg(topic, data)
//}

//func (c *Connection) BroadcastAMOPMsg(topic string, data []byte) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.broadcastAMOPMsg(topic, data)
//}

//func (c *Connection) SendAMOPPrivateMsg(topic string, data []byte) ([]byte, error) {
//	hc := c.writeConn.(*channelSession)
//	return hc.sendAMOPPrivateMsg(topic, data)
//}

//func (c *Connection) BroadcastAMOPPrivateMsg(topic string, data []byte) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.broadcastAMOPPrivateMsg(topic, data)
//}

//func (c *Connection) SubscribeBlockNumberNotify(groupID uint64, handler func(int64)) error {
//	hc := c.writeConn.(*channelSession)
//	return hc.subscribeBlockNumberNotify(groupID, handler)
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

// IsHTTP returns whether is HTTP
func (c *Connection) IsHTTP() bool {
	return c.isHTTP
}

// GetBlockNumber returns BlockLimit
func (c *Connection) GetBlockNumber() int64 {
	hc, ok := c.writeConn.(*channelSession)
	if !ok {
		return 0
	}
	return hc.nodeInfo.blockNumber
}
