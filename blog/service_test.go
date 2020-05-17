package blog

import (
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/ex"
)

func Test_service(t *testing.T) {
	s := NewService("templates")
	assert := asserter.New(t)
	exp := assert().ResponseFrom(s)

	bad := &Article{}
	exp.StatusCode(400, "POST", "/articles", bad.Reader())
	A := &Article{Title: "first", Content: "first"}
	B := &Article{Title: "second", Content: "second"}
	exp.StatusCode(201, "POST", "/articles", A.Reader())
	exp.StatusCode(201, "POST", "/articles", B.Reader())
	assert(len(s.blog) == 2).Fatal("articles not saved")
	exp.StatusCode(200, "GET", "/articles", nil)
	exp.StatusCode(200, "GET", "/", nil)
	exp.Contains("first", "GET", "/", nil)
	exp.StatusCode(404, "DELETE", "/articles/nosuch")
	exp.StatusCode(204, "DELETE", "/articles/first")
	assert(len(s.blog) == 1).Fatal("article not deleted")
}

func ExampleService_ServeHTTP() {
	s := NewService("templates")
	s.ServeHTTP(ex.JsonOf(http.NewRequest("GET", "/articles", nil)))
	// output:
	// []
}
