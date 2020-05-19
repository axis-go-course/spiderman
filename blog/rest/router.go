package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/axis-go-course/spiderman/blog"
	"github.com/gorilla/mux"
)

func NewRouter(s *blog.Service) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/articles", CreateArticle(s.Blog)).Methods("POST")
	r.Handle("/articles", ReadArticles(s.Blog)).Methods("GET")
	r.Handle("/articles/{title}", DeleteArticle(s.Blog)).Methods("DELETE")
	r.Handle("/", blog.UserInterface(s))

	// middlewares
	r.Use(logAllRequests)
	return r
}

func logAllRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func CreateArticle(b blog.Blog) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var v blog.Article
		json.NewDecoder(r.Body).Decode(&v)
		if err := b.SaveArticle(&v); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		io.Copy(w, v.Reader())
	}
}

func ReadArticles(b blog.Blog) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*blog.Article, 5)
		n := b.LoadArticles(result)
		json.NewEncoder(w).Encode(result[:n])
	}
}

func DeleteArticle(b blog.Blog) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := b.DeleteArticle(vars["title"]); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
