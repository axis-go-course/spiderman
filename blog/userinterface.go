package blog

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (s *Service) UserInterface(dir string) http.HandlerFunc {
	pattern := filepath.Join(dir, "*.tmpl")
	t := template.Must(template.ParseGlob(pattern))
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*Article, 10)
		n := s.blog.LoadArticles(result)
		t.ExecuteTemplate(w, "index", struct {
			BlogName string
			Articles []*Article
		}{
			"Spidermans blog",
			result[:n],
		})
	}
}
