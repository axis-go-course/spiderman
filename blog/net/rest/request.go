package rest

import (
	"net/http"
	"net/url"
)

type httpReq struct {
	r *http.Request
}

func (a *httpReq) Header() http.Header { return a.r.Header }
func (a *httpReq) Query() url.Values   { return a.r.URL.Query() }

func (a *httpReq) Read(p []byte) (int, error) {
	if a.r.Body == nil {
		return 0, nil
	}
	return a.r.Body.Read(p)
}
