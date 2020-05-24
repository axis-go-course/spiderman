package rest

import (
	"net/http"
	"net/url"
)

type Request struct {
	*http.Request
}

func (r *Request) Header() http.Header { return r.Request.Header }
func (r *Request) Query() url.Values   { return r.URL.Query() }

func (r *Request) Read(p []byte) (int, error) {
	if r.Body == nil {
		return 0, nil
	}
	return r.Body.Read(p)
}
