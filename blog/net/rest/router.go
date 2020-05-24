// Package rest exposes the blog network API via REST handlers
package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/axis-go-course/spiderman/blog"
	"github.com/gorilla/mux"
)

func NewRouter(page blog.Page, ui http.Handler) *mux.Router {
	r := mux.NewRouter()
	r.Handle("/articles", CreateArticle(page)).Methods("POST")
	r.Handle("/articles", ReadArticles(page)).Methods("GET")
	r.Handle("/articles/{title}", DeleteArticle(page)).Methods("DELETE")
	r.Handle("/", ui)

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

func CreateArticle(page blog.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var article blog.Article
		json.NewDecoder(r.Body).Decode(&article)
		author := blog.NewAuthor("anonymous")
		if err := author.WriteArticle(page, &article); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		io.Copy(w, article.Reader())
	}
}

func ReadArticles(page blog.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*blog.Article, 5)
		n := page.LoadArticles(result)
		json.NewEncoder(w).Encode(result[:n])
	}
}

func DeleteArticle(page blog.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		author := blog.NewAuthor("anonymous")
		article := &blog.Article{Title: vars["title"]}
		if err := author.EraseArticle(page, article); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
