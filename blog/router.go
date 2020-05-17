package blog

import "net/http"

func NewRouter() *Router {
	return &Router{}
}

type Router struct{}

func (r *Router) ServeHTTP(w http.ResponseWriter, q *http.Request) {

}
