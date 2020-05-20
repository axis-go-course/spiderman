package ui

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/axis-go-course/spiderman/blog"
)

func UserInterface(s *blog.Service) http.HandlerFunc {
	pattern := filepath.Join(s.TemplatesDir, "*.tmpl")
	t := template.Must(template.ParseGlob(pattern))
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*blog.Article, 10)
		n := s.Blog.LoadArticles(result)
		t.ExecuteTemplate(w, "index", struct {
			BlogName string
			Articles []*blog.Article
		}{
			"Spidermans blog",
			result[:n],
		})
	}
}
