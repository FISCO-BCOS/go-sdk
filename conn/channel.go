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

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	tls "github.com/FISCO-BCOS/crypto/tls"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/google/uuid"
)

const (
	maxTopicLength      = 254
	messageHeaderLength = 42
	protocolVersion     = 3
	clientType          = "Go-SDK"
)

type nodeInfo struct {
	blockNumber       int64
	Protocol          int32  `json:"protocol"`
	CompatibleVersion string `json:"nodeVersion"`
}

type channelSession struct {
	// groupID   uint
	c         *tls.Conn
	mu        sync.RWMutex
	responses map[string]*channelResponse
	// receiptsMutex sync.Mutex
	receipts      map[string]*channelResponse
	topicMu       sync.RWMutex
	topicHandlers map[string]func(*topicData)
	buf           []byte
	nodeInfo      nodeInfo
	closeOnce     sync.Once
	closed        chan interface{}
}

const (
	// channel messages types
	rpcMessage             = 0x12   // channel rpc request
	clientHeartbeat        = 0x13   // Heartbeat for sdk
	clientHandshake        = 0x14   // type for hand shake
	clientRegisterEventLog = 0x15   // type for event log filter register request and response
	amopPushRandom         = 0x30   // type for request from sdk
	amopResponse           = 0x31   // type for response to sdk
	amopSubscribeTopics    = 0x32   // type for topic request
	amopMultiCast          = 0x35   // type for mult broadcast
	requestTopicCert       = 0x38   // type for update status
	transactionNotify      = 0x1000 // type for  transaction notify
	blockNotify            = 0x1001 // type for  block notify
	eventLogPush           = 0x1002 // type for event log push
)

type topicData struct {
	length uint8
	topic  string
	data   []byte
}

type channelMessage struct {
	length    uint32
	typeN     uint16
	uuid      string
	errorCode int32
	body      []byte
}

type handshakeRequest struct {
	MinimumSupport int32  `json:"minimumSupport"`
	MaximumSupport int32  `json:"maximumSupport"`
	ClientType     string `json:"clientType"`
}

type handshakeResponse struct {
	Protocol    int32  `json:"protocol"`
	NodeVersion string `json:"nodeVersion"`
}

type channelResponse struct {
	Message *channelMessage
	Notify  chan interface{}
}

func newChannelMessage(msgType uint16, body []byte) (*channelMessage, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Fatal("newChannelMessage error:", err)
		return nil, err
	}
	idString := strings.ReplaceAll(id.String(), "-", "")
	// var idByte [32]byte
	// copy(idByte[:], idString[:32])
	msg := &channelMessage{length: uint32(messageHeaderLength + len(body)), typeN: msgType,
		errorCode: 0, uuid: idString, body: body}
	return msg, nil
}

func newTopicMessage(t string, data []byte, msgType uint16) (*channelMessage, error) {
	if len(t) > maxTopicLength {
		return nil, fmt.Errorf("topic length exceeds 255")
	}
	topic := &topicData{length: uint8(len(t)) + 1, topic: t, data: data}
	mesgData := topic.Encode()
	return newChannelMessage(msgType, mesgData)
}

func (t *topicData) Encode() []byte {
	var raw []byte
	buf := bytes.NewBuffer(raw)
	err := binary.Write(buf, binary.LittleEndian, t.length)
	if err != nil {
		log.Fatal("encode length error:", err)
	}
	err = binary.Write(buf, binary.LittleEndian, t.topic)
	if err != nil {
		log.Fatal("encode type error:", err)
	}
	err = binary.Write(buf, binary.LittleEndian, t.data)
	if err != nil {
		log.Fatal("encode uuid error:", err)
	}
	return buf.Bytes()
}

func (msg *channelMessage) Encode() []byte {
	var raw []byte
	buf := bytes.NewBuffer(raw)
	err := binary.Write(buf, binary.BigEndian, msg.length)
	if err != nil {
		log.Fatal("encode length error:", err)
	}
	err = binary.Write(buf, binary.BigEndian, msg.typeN)
	if err != nil {
		log.Fatal("encode type error:", err)
	}
	err = binary.Write(buf, binary.LittleEndian, []byte(msg.uuid))
	if err != nil {
		log.Fatal("encode uuid error:", err)
	}
	err = binary.Write(buf, binary.BigEndian, msg.errorCode)
	if err != nil {
		log.Fatal("encode ErrorCode error:", err)
	}
	err = binary.Write(buf, binary.LittleEndian, msg.body)
	if err != nil {
		log.Fatal("encode Body error:", err)
	}
	if uint32(buf.Len()) != msg.length {
		fmt.Printf("%d != %d\n, buf is %v", buf.Len(), msg.length, buf.String())
		log.Fatal("encode error length error:", err)
	}
	return buf.Bytes()
}

func decodeChannelMessage(raw []byte) (*channelMessage, error) {
	buf := bytes.NewReader(raw)
	result := new(channelMessage)
	err := binary.Read(buf, binary.BigEndian, &result.length)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	if uint32(len(raw)) < result.length {
		return nil, errors.New("uncomplete message")
	}
	err = binary.Read(buf, binary.BigEndian, &result.typeN)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	var uuid [32]byte
	err = binary.Read(buf, binary.LittleEndian, &uuid)
	if err != nil {
		// log.Fatal("encode error:", err)
		fmt.Println("binary.Read failed:", err)
	}
	result.uuid = string(uuid[:])

	err = binary.Read(buf, binary.BigEndian, &result.errorCode)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	dataLength := result.length - messageHeaderLength
	result.body = make([]byte, dataLength, dataLength)
	err = binary.Read(buf, binary.BigEndian, &result.body)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return result, nil
}

func decodeTopic(raw []byte) (*topicData, error) {
	buf := bytes.NewReader(raw)
	result := new(topicData)
	err := binary.Read(buf, binary.LittleEndian, &result.length)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	topic := make([]byte, result.length-1, result.length-1)
	err = binary.Read(buf, binary.LittleEndian, &topic)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	result.topic = string(topic)
	dataLength := len(raw) - int(result.length)
	result.data = make([]byte, dataLength, dataLength)
	err = binary.Read(buf, binary.LittleEndian, &result.data)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	return result, nil
}

// channelCon n is treated specially by Connection.
func (hc *channelSession) Write(context.Context, interface{}) error {
	panic("Write called on channelSession")
}

func (hc *channelSession) RemoteAddr() string {
	return hc.c.RemoteAddr().String()
}

func (hc *channelSession) Read() ([]*jsonrpcMessage, bool, error) {
	<-hc.closed
	return nil, false, io.EOF
}

func (hc *channelSession) Close() {
	hc.closeOnce.Do(func() { close(hc.closed) })
}

func (hc *channelSession) Closed() <-chan interface{} {
	return hc.closed
}

// ChannelTimeouts represents the configuration params for the Channel RPC server.
type ChannelTimeouts struct {
	// ReadTimeout is the maximum duration for reading the entire
	// request, including the body.
	//
	// Because ReadTimeout does not let Handlers make per-request
	// decisions on each request body's acceptable deadline or
	// upload rate, most users will prefer to use
	// ReadHeaderTimeout. It is valid to use them both.
	ReadTimeout time.Duration

	// WriteTimeout is the maximum duration before timing out
	// writes of the response. It is reset whenever a new
	// request's header is read. Like ReadTimeout, it does not
	// let Handlers make decisions on a per-request basis.
	WriteTimeout time.Duration

	// IdleTimeout is the maximum amount of time to wait for the
	// next request when keep-alives are enabled. If IdleTimeout
	// is zero, the value of ReadTimeout is used. If both are
	// zero, ReadHeaderTimeout is used.
	IdleTimeout time.Duration
}

// DefaultChannelTimeouts represents the default timeout values used if further
// configuration is not provided.
var DefaultChannelTimeouts = ChannelTimeouts{
	ReadTimeout:  30 * time.Second,
	WriteTimeout: 30 * time.Second,
	IdleTimeout:  120 * time.Second,
}

// DialChannelWithClient creates a new RPC client that connects to an RPC server over Channel
// using the provided Channel Client.
func DialChannelWithClient(endpoint string, config *tls.Config, groupID int) (*Connection, error) {
	initctx := context.Background()
	return newClient(initctx, func(context.Context) (ServerCodec, error) {
		conn, err := tls.Dial("tcp", endpoint, config)
		if err != nil {
			return nil, err
		}
		ch := &channelSession{c: conn, responses: make(map[string]*channelResponse),
			receipts: make(map[string]*channelResponse), topicHandlers: make(map[string]func(*topicData)),
			nodeInfo: nodeInfo{blockNumber: 0, Protocol: 1}, closed: make(chan interface{})}
		go ch.processMessages()
		if err = ch.handshakeChannel(); err != nil {
			fmt.Printf("handshake channel protocol failed, use default protocol version")
		}
		if err = ch.SubscribeTopic("_block_notify_"+strconv.Itoa(groupID), nil); err != nil {
			return nil, fmt.Errorf("subscribe block nofity failed")
		}
		return ch, nil
	})
}

func (c *Connection) sendRPCRequest(ctx context.Context, op *requestOp, msg interface{}) error {
	hc := c.writeConn.(*channelSession)
	rpcMsg := msg.(*jsonrpcMessage)
	if rpcMsg.Method == "sendRawTransaction" {
		respBody, err := hc.sendTransaction(ctx, msg)
		if err != nil {
			return fmt.Errorf("sendTransaction %v", err)
		}
		rpcResp := new(jsonrpcMessage)
		rpcResp.Result = respBody
		op.resp <- rpcResp
	} else {
		respBody, err := hc.doRPCRequest(ctx, msg)
		if respBody != nil {
			defer respBody.Close()
		}

		if err != nil {
			if respBody != nil {
				buf := new(bytes.Buffer)
				if _, err2 := buf.ReadFrom(respBody); err2 == nil {
					return fmt.Errorf("%v %v", err, buf.String())
				}
			}
			return err
		}
		var respmsg jsonrpcMessage
		if err := json.NewDecoder(respBody).Decode(&respmsg); err != nil {
			return err
		}
		op.resp <- &respmsg
	}
	return nil
}

func (c *Connection) sendBatchChannel(ctx context.Context, op *requestOp, msgs []*jsonrpcMessage) error {
	hc := c.writeConn.(*channelSession)
	respBody, err := hc.doRPCRequest(ctx, msgs)
	if err != nil {
		return err
	}
	defer respBody.Close()
	var respmsgs []jsonrpcMessage
	if err := json.NewDecoder(respBody).Decode(&respmsgs); err != nil {
		return err
	}
	for i := 0; i < len(respmsgs); i++ {
		op.resp <- &respmsgs[i]
	}
	return nil
}

func (hc *channelSession) doRPCRequest(ctx context.Context, msg interface{}) (io.ReadCloser, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	var rpcMsg *channelMessage
	rpcMsg, err = newChannelMessage(rpcMessage, body)
	if err != nil {
		return nil, err
	}
	msgBytes := rpcMsg.Encode()

	_, err = hc.c.Write(msgBytes)
	if err != nil {
		return nil, err
	}
	response := &channelResponse{Message: nil, Notify: make(chan interface{})}
	hc.mu.Lock()
	hc.responses[rpcMsg.uuid] = response
	hc.mu.Unlock()

	<-response.Notify
	hc.mu.Lock()
	response = hc.responses[rpcMsg.uuid]
	delete(hc.responses, rpcMsg.uuid)
	hc.mu.Unlock()
	if response.Message.errorCode != 0 {
		return nil, errors.New("response error:" + string(response.Message.errorCode))
	}
	return ioutil.NopCloser(bytes.NewReader(response.Message.body)), nil
}

func (hc *channelSession) sendTransaction(ctx context.Context, msg interface{}) ([]byte, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	var rpcMsg *channelMessage
	rpcMsg, err = newChannelMessage(rpcMessage, body)
	if err != nil {
		return nil, err
	}

	response := &channelResponse{Message: nil, Notify: make(chan interface{})}
	receipt := &channelResponse{Message: nil, Notify: make(chan interface{})}
	hc.mu.Lock()
	hc.responses[rpcMsg.uuid] = response
	hc.receipts[rpcMsg.uuid] = receipt
	hc.mu.Unlock()
	defer func() {
		hc.mu.Lock()
		delete(hc.responses, rpcMsg.uuid)
		delete(hc.receipts, rpcMsg.uuid)
		hc.mu.Unlock()
	}()
	msgBytes := rpcMsg.Encode()
	_, err = hc.c.Write(msgBytes)
	if err != nil {
		return nil, err
	}
	<-response.Notify

	hc.mu.Lock()
	response = hc.responses[rpcMsg.uuid]
	delete(hc.responses, rpcMsg.uuid)
	hc.mu.Unlock()
	if response.Message.errorCode != 0 {
		return nil, errors.New("response error:" + string(response.Message.errorCode))
	}
	var respmsg jsonrpcMessage
	if err := json.NewDecoder(bytes.NewReader(response.Message.body)).Decode(&respmsg); err != nil {
		return nil, err
	}
	if respmsg.Error != nil {
		return nil, fmt.Errorf("send transaction error, code=%d, message=%s", respmsg.Error.Code, respmsg.Error.Message)
	}
	// fmt.Printf("sendTransaction reveived response,seq:%s message:%s\n ", rpcMsg.uuid, respmsg.Result)

	<-receipt.Notify

	hc.mu.RLock()
	receipt = hc.receipts[rpcMsg.uuid]
	hc.mu.RUnlock()
	if receipt.Message.errorCode != 0 {
		return nil, errors.New("response error:" + string(receipt.Message.errorCode))
	}
	var transactionReceipt types.Receipt
	if err := json.Unmarshal(receipt.Message.body, &transactionReceipt); err != nil {
		return nil, fmt.Errorf("parse receipt error %w", err)
	}
	if transactionReceipt.Status != "0x0" {
		return nil, fmt.Errorf("receipt error code:%s", transactionReceipt.Status)
	}
	// fmt.Printf("sendTransaction reveived transactionReceipt:%+v\n ", transactionReceipt)
	return receipt.Message.body, nil
}

func (hc *channelSession) doRequestNoResponse(msg *channelMessage) error {
	msgBytes := msg.Encode()
	_, err := hc.c.Write(msgBytes)
	if err != nil {
		return err
	}
	return nil
}

func (hc *channelSession) doRequest(msg *channelMessage) (*channelMessage, error) {
	msgBytes := msg.Encode()
	response := &channelResponse{Message: nil, Notify: make(chan interface{})}
	hc.mu.Lock()
	hc.responses[msg.uuid] = response
	hc.mu.Unlock()
	_, err := hc.c.Write(msgBytes)
	if err != nil {
		return nil, err
	}
	defer func() {
		hc.mu.Lock()
		delete(hc.responses, msg.uuid)
		hc.mu.Unlock()
	}()

	<-response.Notify
	hc.mu.Lock()
	response = hc.responses[msg.uuid]
	hc.mu.Unlock()
	if response.Message.errorCode != 0 {
		return nil, errors.New("response error:" + string(response.Message.errorCode))
	}
	return response.Message, nil
}

func (hc *channelSession) handshakeChannel() error {
	handshakeBody := handshakeRequest{MinimumSupport: 1, MaximumSupport: protocolVersion, ClientType: clientType}
	body, err := json.Marshal(handshakeBody)
	if err != nil {
		return fmt.Errorf("encode handshake request failed %w", err)
	}
	var msg, response *channelMessage
	msg, err = newChannelMessage(clientHandshake, body)
	response, err = hc.doRequest(msg)
	var info nodeInfo
	if err = json.Unmarshal(response.body, &info); err != nil {
		return fmt.Errorf("parse handshake channel protocol response failed %w", err)
	}
	hc.nodeInfo = info
	// fmt.Printf("node info:%+v", info)
	return nil
}

func (hc *channelSession) SubscribeTopic(topic string, handler func(*topicData)) error {
	if _, ok := hc.topicHandlers[topic]; ok {
		return errors.New("already subscribed to topic " + topic)
	}
	hc.topicMu.Lock()
	hc.topicHandlers[topic] = handler
	hc.topicMu.Unlock()

	keys := make([]string, 0, len(hc.topicHandlers))
	for k := range hc.topicHandlers {
		keys = append(keys, k)
	}
	data, err := json.Marshal(keys)
	if err != nil {
		hc.topicMu.Lock()
		delete(hc.topicHandlers, topic)
		hc.topicMu.Unlock()
		return errors.New("marshal topics failed")
	}
	msg, err := newChannelMessage(amopSubscribeTopics, data)
	return hc.doRequestNoResponse(msg)
}

func (hc *channelSession) UnsubscribeTopic(topic string) error {
	hc.topicMu.Lock()
	delete(hc.topicHandlers, topic)
	hc.topicMu.Unlock()

	keys := make([]string, 0, len(hc.topicHandlers))
	for k := range hc.topicHandlers {
		keys = append(keys, k)
	}
	data, err := json.Marshal(keys)
	if err != nil {
		return errors.New("marshal topics failed")
	}
	msg, err := newChannelMessage(amopSubscribeTopics, data)
	return hc.doRequestNoResponse(msg)
}

func (hc *channelSession) PushTopicDataRandom(topic string, data []byte) error {
	msg, err := newTopicMessage(topic, data, amopPushRandom)
	if err != nil {
		return err
	}
	return hc.doRequestNoResponse(msg)
}

func (hc *channelSession) PushTopicDataToALL(topic string, data []byte) error {
	msg, err := newTopicMessage(topic, data, amopMultiCast)
	if err != nil {
		return err
	}
	return hc.doRequestNoResponse(msg)
}

func (hc *channelSession) processTopicResponse(msg *channelMessage) error {
	hc.processTopicResponse(msg)
	topic, err := decodeTopic(msg.body)
	if err != nil {
		fmt.Printf("decode topic failed: %+v", msg)
		return err
	}
	hc.topicMu.RLock()
	handler, ok := hc.topicHandlers[topic.topic]
	hc.topicMu.RUnlock()
	if !ok {
		return fmt.Errorf("unsubscribe topic %s", topic.topic)
	}
	handler(topic)
	return nil
}

func (hc *channelSession) processMessages() {
	for {
		select {
		case <-hc.closed:
			return
		default:
			receiveBuf := make([]byte, 4096)
			b, err := hc.c.Read(receiveBuf)
			if err != nil {
				// fmt.Printf("channel Read error:%v", err)
				hc.Close()
			}
			hc.buf = append(hc.buf, receiveBuf[:b]...)
			msg, err := decodeChannelMessage(hc.buf)
			if err != nil {
				// fmt.Printf("decodeChannelMessage error:%v", err)
				continue
			}
			// fmt.Printf("message %+v\n", msg)
			hc.buf = hc.buf[msg.length:]
			hc.mu.Lock()
			if response, ok := hc.responses[msg.uuid]; ok {
				response.Message = msg
				response.Notify <- struct{}{}
			}
			hc.mu.Unlock()
			switch msg.typeN {
			case rpcMessage, clientHandshake:
				// fmt.Printf("response type:%d seq:%s, msg:%s", msg.typeN, msg.uuid, string(msg.body))
			case transactionNotify:
				// fmt.Printf("transaction notify:%s", string(msg.body))
				hc.mu.Lock()
				if receipt, ok := hc.receipts[msg.uuid]; ok {
					receipt.Message = msg
					receipt.Notify <- struct{}{}
				} else {
					fmt.Printf("error %+v", receipt)
				}
				hc.mu.Unlock()
			case blockNotify:
				hc.updateBlockNumber(msg)
			case amopResponse:
				// response of pushRandom and broadcast
				err := hc.processTopicResponse(msg)
				fmt.Printf("response type:%d seq:%s, msg:%s, err:%v", msg.typeN, msg.uuid, string(msg.body), err)
			default:
				fmt.Printf("unknown message type:%d, msg:%+v", msg.typeN, msg)
			}
		}
	}
}

func (hc *channelSession) updateBlockNumber(msg *channelMessage) {
	var blockNumber int64
	topic, err := decodeTopic(msg.body)
	if hc.nodeInfo.Protocol == 1 {
		response := strings.Split(string(topic.data), ",")
		blockNumber, err = strconv.ParseInt(response[1], 10, 32)
		if err != nil {
			fmt.Print("v1 block notify parse blockNumber failed")
			return
		}
	} else {
		var notify struct {
			GroupID     uint  `json:"groupID"`
			BlockNumber int64 `json:"blockNumber"`
		}
		err = json.Unmarshal(topic.data, &notify)
		if err != nil {
			fmt.Print("block notify parse blockNumber failed")
			return
		}
		blockNumber = notify.BlockNumber
	}
	// fmt.Printf("blockNumber updated %d -> %d", hc.nodeInfo.blockNumber, blockNumber)
	hc.nodeInfo.blockNumber = blockNumber
}

// channelServerConn turns a Channel connection into a Conn.
type channelServerConn struct {
	io.Reader
	io.Writer
	r *http.Request
}

func newChannelServerConn(r *http.Request, w http.ResponseWriter) ServerCodec {
	body := io.LimitReader(r.Body, maxRequestContentLength)
	conn := &channelServerConn{Reader: body, Writer: w, r: r}
	return NewJSONCodec(conn)
}

// Close does nothing and always returns nil.
func (t *channelServerConn) Close() error { return nil }

// RemoteAddr returns the peer address of the underlying connection.
func (t *channelServerConn) RemoteAddr() string {
	return t.r.RemoteAddr
}

// SetWriteDeadline does nothing and always returns nil.
func (t *channelServerConn) SetWriteDeadline(time.Time) error { return nil }
