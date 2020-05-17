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
	router := NewService()
	assert := asserter.New(t)
	exp := assert().ResponseFrom(router)
	exp.StatusCode(200, "GET", "/articles", nil)
	exp.BodyIs("{}\n", "GET", "/articles", nil)

	article := &Article{Title: "hello world", Content: "more here"}
	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(article)
		w.Close()
	}()
	exp.StatusCode(201, "POST", "/articles", r)
}

func ExampleRouter_ServeHTTP() {
	router := NewService()
	router.ServeHTTP(ex.BodyOf(http.NewRequest("GET", "/articles", nil)))
	// output:
	// {}
}
