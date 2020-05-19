package blog

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func UserInterface(s *Service) http.HandlerFunc {
	pattern := filepath.Join(s.TemplatesDir, "*.tmpl")
	t := template.Must(template.ParseGlob(pattern))
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*Article, 10)
		n := s.Blog.LoadArticles(result)
		t.ExecuteTemplate(w, "index", struct {
			BlogName string
			Articles []*Article
		}{
			"Spidermans blog",
			result[:n],
		})
	}
}
