package rest

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func (w *Response) SetHeader(key, val string) {
	w.Header().Set(key, val)
}
