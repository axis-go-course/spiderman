package net

import (
	"io"
	"net/http"
	"net/url"
)

type Request interface {
	io.Reader // the body
	Query() url.Values
	Header() http.Header
}
