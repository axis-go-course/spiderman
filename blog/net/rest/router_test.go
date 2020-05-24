package rest

import (
	"net/http"
	"testing"

	"github.com/axis-go-course/spiderman/blog"
	"github.com/axis-go-course/spiderman/blog/ui"
	"github.com/gregoryv/asserter"
	"github.com/gregoryv/ex"
)

func Test_service(t *testing.T) {
	page := blog.NewPage()
	ui, err := ui.UserInterface("../../ui/templates", page)
	if err != nil {
		t.Fatal(err)
	}
	r := NewRouter(page, ui)
	assert := asserter.New(t)
	exp := assert().ResponseFrom(r)

	bad := &blog.Article{}
	exp.StatusCode(400, "POST", "/articles", bad.Reader())
	A := &blog.Article{Title: "first", Content: "first"}
	B := &blog.Article{Title: "second", Content: "second"}
	exp.StatusCode(201, "POST", "/articles", A.Reader())
	exp.StatusCode(201, "POST", "/articles", B.Reader())
	assert(len(page) == 2).Fatal("articles not saved")
	exp.StatusCode(200, "GET", "/articles", nil)
	exp.StatusCode(200, "GET", "/", nil)
	exp.Contains("first", "GET", "/", nil)
	exp.StatusCode(404, "DELETE", "/articles/nosuch")
	exp.StatusCode(204, "DELETE", "/articles/first")
	assert(len(page) == 1).Fatal("article not deleted")
}

func ExampleNewRouter_getArticles() {
	page := blog.NewPage()
	r := NewRouter(page, nil)
	r.ServeHTTP(ex.JsonOf(http.NewRequest("GET", "/articles", nil)))
	// output:
	// []
}
