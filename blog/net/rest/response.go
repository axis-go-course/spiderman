package rest

import (
	"net/http"
)

type httpResp struct {
	http.ResponseWriter
}

func (w *httpResp) SetHeader(key, val string) {
	w.Header().Set(key, val)
}
