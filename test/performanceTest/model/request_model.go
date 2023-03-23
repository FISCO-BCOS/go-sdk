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
	// HTTPOk request succeed
	HTTPOk = 200
	// RequestErr request error
	RequestErr = 509
	// ParseError
	ParseError = 510
)

const (
	FormTypeKvTable = "kvTable"
	FormParallelOk  = "parallelOk"
)

// Verify 验证器
type Verify interface {
	GetCode() int
	GetResult() bool // Return success or not
}

// Request request data
type Request struct {
	URL       string            // URL
	Form      string            // http/webSocket/tcp
	Method    string            // method GET/POST/PUT
	Headers   map[string]string // Headers
	Body      string            // body
	Verify    string
	Timeout   time.Duration
	Debug     bool
	MaxCon    int
	HTTP2     bool
	Keepalive bool
	Code      int // status code
}

// GetBody
func (r *Request) GetBody() (body io.Reader) {
	return strings.NewReader(r.Body)
}

// getVerifyKey
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

// RequestResults
type RequestResults struct {
	ID            string // message ID
	ChanID        uint64 // message ID
	Time          uint64 // request time nanosecond
	IsSucceed     bool   // success ot not
	ErrCode       int    // error code
	ReceivedBytes int64
}

// SetID Set Request ID
func (r *RequestResults) SetID(chanID uint64, number uint64) {
	id := fmt.Sprintf("%d_%d", chanID, number)
	r.ID = id
	r.ChanID = chanID
}
