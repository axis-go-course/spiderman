package net

import (
	"bytes"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Returns a lambdaResp with default status code 200 ok
func NewLambdaResp() *lambdaResp {
	resp := lambdaResp{}
	resp.StatusCode = http.StatusOK
	return &resp
}

type lambdaResp struct {
	events.APIGatewayProxyResponse
	buf bytes.Buffer
}

func (w *lambdaResp) SetHeader(key, val string) {
	if w.Headers == nil {
		w.Headers = make(map[string]string)
	}
	w.Headers[key] = val
}

func (w *lambdaResp) Write(data []byte) (int, error) {
	n, err := w.buf.Write(data)
	w.Body = w.buf.String()
	return n, err
}

func (w *lambdaResp) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
}
