package blog

import (
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/ex"
)

func xTestRouter(t *testing.T) {
	s := NewService()
	A := &Article{Title: "first", Content: "first"}
	B := &Article{Title: "second", Content: "second"}
	assert := asserter.New(t)
	exp := assert().ResponseFrom(s)
	exp.StatusCode(201, "POST", "/articles", A.Reader())
	exp.StatusCode(201, "POST", "/articles", B.Reader())
	exp.StatusCode(200, "GET", "/articles", nil)
}

func xExampleRouter_ServeHTTP() {
	s := NewService()
	s.ServeHTTP(ex.JsonOf(http.NewRequest("GET", "/articles", nil)))
	// output:
	// []
}
