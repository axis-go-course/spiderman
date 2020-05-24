package net

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

func NewRequest(r *events.APIGatewayProxyRequest) *Request {
	return &Request{r: r, bodyReader: strings.NewReader(r.Body)}
}

type Request struct {
	r          *events.APIGatewayProxyRequest
	bodyReader io.Reader
}

func (a *Request) Read(p []byte) (int, error) {
	return a.bodyReader.Read(p)
}

func (a *Request) Query() url.Values {
	v := url.Values{}
	for k, in := range a.r.MultiValueQueryStringParameters {
		for _, inval := range in {
			v.Add(k, inval)
		}
	}
	for k, in := range a.r.QueryStringParameters {
		v.Add(k, in)
	}
	return v
}

func (a *Request) Header() http.Header {
	if a.r.MultiValueHeaders == nil {
		a.r.MultiValueHeaders = make(http.Header)
	}
	for k, v := range a.r.Headers {
		a.r.MultiValueHeaders[k] = []string{v}
	}
	return a.r.MultiValueHeaders
}
