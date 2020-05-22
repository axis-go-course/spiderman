package ui

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/axis-go-course/spiderman/blog"
)

func UserInterface(tmplDir string, page blog.Page) (http.HandlerFunc, error) {
	pattern := filepath.Join(tmplDir, "*.tmpl")
	t, err := template.ParseGlob(pattern)
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*blog.Article, 10)
		n := page.LoadArticles(result)
		t.ExecuteTemplate(w, "index", struct {
			BlogName string
			Articles []*blog.Article
		}{
			"Spidermans blog",
			result[:n],
		})
	}, err
}
