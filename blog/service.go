package blog

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewService() *Service {
	r := mux.NewRouter()
	s := &Service{
		router: r,
		blog:   NewBlog(),
	}
	r.Handle("/articles", s.ReadArticles()).Methods("GET")
	r.Handle("/articles", s.CreateArticle()).Methods("POST")
	return s
}

type Service struct {
	router *mux.Router
	blog   Blog
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Service) CreateArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func (s *Service) ReadArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*Article, 5)
		n := s.blog.LoadArticles(result)
		json.NewEncoder(w).Encode(result[:n])
	}
}
