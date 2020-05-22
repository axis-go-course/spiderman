package ui

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/axis-go-course/spiderman/blog"
)

func UserInterface(tmplDir string, page blog.Page) http.HandlerFunc {
	pattern := filepath.Join(tmplDir, "*.tmpl")
	t := template.Must(template.ParseGlob(pattern))
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
	}
}
