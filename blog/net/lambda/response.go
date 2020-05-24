package net

import (
	"bytes"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// NewResponse returns a response with default status code 200 ok
func NewResponse() *Response {
	resp := Response{}
	resp.StatusCode = http.StatusOK
	return &resp
}

type Response struct {
	events.APIGatewayProxyResponse
	buf bytes.Buffer
}

func (w *Response) SetHeader(key, val string) {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	w.Headers[key] = val
}

func (w *Response) Write(data []byte) (int, error) {
	n, err := w.buf.Write(data)
	w.Body = w.buf.String()
	return n, err
}

func (w *Response) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}
