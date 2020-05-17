package blog

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/ex"
)

func TestRouter(t *testing.T) {
	A := &Article{Title: "first", Content: "first"}
	B := &Article{Title: "second", Content: "second"}
	all := []*Article{A, B}
	s := NewService()
	s.blog.SaveArticle(all...)

	assert := asserter.New(t)
	exp := assert().ResponseFrom(s)
	exp.StatusCode(200, "GET", "/articles", nil)

	article := &Article{Title: "hello world", Content: "more here"}
	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(article)
		w.Close()
	}()
	exp.StatusCode(201, "POST", "/articles", r)
}

func ExampleRouter_ServeHTTP() {
	s := NewService()
	s.ServeHTTP(ex.JsonOf(http.NewRequest("GET", "/articles", nil)))
	// output:
	// []
}
