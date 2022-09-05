// Package model 请求数据模型package model
package model

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// 返回 code 码
const (
	// HTTPOk 请求成功
	HTTPOk = 200
	// RequestErr 请求错误
	RequestErr = 509
	// ParseError 解析错误
	ParseError = 510 // 解析错误
)

// 支持协议
const (
	FormTypeHTTP = "http"
	FormTypeWebSocket = "webSocket"
	// FormTypeGRPC grpc 协议
	FormTypeGRPC    = "grpc"
	FormTypeKvTable = "kvTable"
	FormParallelOk  = "parallelOk"
	FormTypeamop    = "amop"
)

// Verify 验证器
type Verify interface {
	GetCode() int    // 有一个方法，返回code为200为成功
	GetResult() bool // 返回是否成功
}


// Request 请求数据
type Request struct {
	URL       string            // URL
	Form      string            // http/webSocket/tcp
	Method    string            // 方法 GET/POST/PUT
	Headers   map[string]string // Headers
	Body      string            // body
	Verify    string            // 验证的方法
	Timeout   time.Duration     // 请求超时时间
	Debug     bool              // 是否开启Debug模式
	MaxCon    int               // 每个连接的请求数
	HTTP2     bool              // 是否使用http2.0
	Keepalive bool              // 是否开启长连接
	Code      int               // 验证的状态码
}

// GetBody 获取请求数据
func (r *Request) GetBody() (body io.Reader) {
	return strings.NewReader(r.Body)
}

// getVerifyKey 获取校验 key
func (r *Request) getVerifyKey() (key string) {
	return fmt.Sprintf("%s.%s", r.Form, r.Verify)
}


func NewRequestByContractType(contractType string, contractMothod string) (request *Request, err error) {
	form := ""
	switch contractType {
	case "kvTableTest":
		form = FormTypeKvTable
	case "parallelOk":
		form = FormParallelOk
	default:
		return nil, errors.New("no this type")
	}
	request = &Request{
		Form:   form,
		Method: contractMothod,
	}
	return
}


// getHeaderValue 获取 header
func getHeaderValue(v string, headers map[string]string) {
	index := strings.Index(v, ":")
	if index < 0 {
		return
	}
	vIndex := index + 1
	if len(v) >= vIndex {
		value := strings.TrimPrefix(v[vIndex:], " ")
		if _, ok := headers[v[:index]]; ok {
			headers[v[:index]] = fmt.Sprintf("%s; %s", headers[v[:index]], value)
		} else {
			headers[v[:index]] = value
		}
	}
}


// RequestResults 请求结果
type RequestResults struct {
	ID            string // 消息ID
	ChanID        uint64 // 消息ID
	Time          uint64 // 请求时间 纳秒
	IsSucceed     bool   // 是否请求成功
	ErrCode       int    // 错误码
	ReceivedBytes int64
}

// SetID 设置请求唯一ID
func (r *RequestResults) SetID(chanID uint64, number uint64) {
	id := fmt.Sprintf("%d_%d", chanID, number)
	r.ID = id
	r.ChanID = chanID
}
