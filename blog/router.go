package blog

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(s *Service) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/articles", CreateArticle(s)).Methods("POST")
	r.Handle("/articles", ReadArticles(s)).Methods("GET")
	r.Handle("/articles/{title}", DeleteArticle(s)).Methods("DELETE")
	r.Handle("/", s.UserInterface(s.templatesDir))
	return r
}

func CreateArticle(s *Service) http.HandlerFunc {
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

func ReadArticles(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*Article, 5)
		n := s.blog.LoadArticles(result)
		json.NewEncoder(w).Encode(result[:n])
	}
}

func DeleteArticle(s *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := s.blog.DeleteArticle(vars["title"]); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
