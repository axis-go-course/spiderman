package blog

import (
	"encoding/json"
	"io"
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
	r.Handle("/articles/{title}", s.DeleteArticle()).Methods("DELETE")
	r.Handle("/", s.UserInterface())
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
		var v Article
		json.NewDecoder(r.Body).Decode(&v)
		if err := s.blog.SaveArticle(&v); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		io.Copy(w, v.Reader())
	}
}

func (s *Service) ReadArticles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*Article, 5)
		n := s.blog.LoadArticles(result)
		json.NewEncoder(w).Encode(result[:n])
	}
}

func (s *Service) DeleteArticle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := s.blog.DeleteArticle(vars["title"]); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
