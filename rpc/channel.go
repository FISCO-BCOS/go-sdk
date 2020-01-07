package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type channelConn struct {
	client    *http.Client
	req       *http.Request
	closeOnce sync.Once
	closed    chan interface{}
}

// channelCon n is treated specially by Connection.
func (hc *channelConn) Write(context.Context, interface{}) error {
	panic("Write called on channelConn")
}

func (hc *channelConn) RemoteAddr() string {
	return hc.req.URL.String()
}

func (hc *channelConn) Read() ([]*jsonrpcMessage, bool, error) {
	<-hc.closed
	return nil, false, io.EOF
}

func (hc *channelConn) Close() {
	hc.closeOnce.Do(func() { close(hc.closed) })
}

func (hc *channelConn) Closed() <-chan interface{} {
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
func DialChannelWithClient(endpoint string, client *http.Client) (*Connection, error) {
	req, err := http.NewRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", contentType)

	initctx := context.Background()
	return newClient(initctx, func(context.Context) (ServerCodec, error) {
		return &channelConn{client: client, req: req, closed: make(chan interface{})}, nil
	})
}

// DialChannel creates a new client that connects to node over tls.
func DialChannel(endpoint string) (*Connection, error) {
	return DialChannelWithClient(endpoint, new(http.Client))
}

func (c *Connection) sendChannel(ctx context.Context, op *requestOp, msg interface{}) error {
	hc := c.writeConn.(*channelConn)
	respBody, err := hc.doRequest(ctx, msg)
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
	return nil
}

func (c *Connection) sendBatchChannel(ctx context.Context, op *requestOp, msgs []*jsonrpcMessage) error {
	hc := c.writeConn.(*channelConn)
	respBody, err := hc.doRequest(ctx, msgs)
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

func (hc *channelConn) doRequest(ctx context.Context, msg interface{}) (io.ReadCloser, error) {
	body, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	req := hc.req.WithContext(ctx)
	req.Body = ioutil.NopCloser(bytes.NewReader(body))
	req.ContentLength = int64(len(body))

	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp.Body, errors.New(resp.Status)
	}
	return resp.Body, nil
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
