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

func ReadArticles(b blog.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := make([]*blog.Article, 5)
		n := b.LoadArticles(result)
		json.NewEncoder(w).Encode(result[:n])
	}
}

func DeleteArticle(b blog.Page) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if err := b.DeleteArticle(vars["title"]); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
