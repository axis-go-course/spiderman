package blog

import (
	"net/http"
	"testing"

	"github.com/gregoryv/asserter"
	"github.com/gregoryv/ex"
)

func TestRouter(t *testing.T) {
	router := NewRouter()
	assert := asserter.New(t)
	exp := assert().ResponseFrom(router)
	exp.StatusCode(200, "GET", "/", nil)
	exp.BodyIs("{}\n", "GET", "/", nil)
}

func ExampleRouter_ServeHTTP() {
	router := NewRouter()
	router.ServeHTTP(ex.BodyOf(http.NewRequest("GET", "/", nil)))
	// output:
	// {}
}
