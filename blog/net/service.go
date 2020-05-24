// Package net implements network API of the spiderman blog.
package net

import "github.com/axis-go-course/spiderman/blog"

type Service struct {
	page *blog.Page
}

func (s *Service) CreateArticle(w Response, r Request) {
	// todo
}

func (s *Service) ReadArticles(w Response, r Request) {
	// todo
}

func (s *Service) DeleteArticle(w Response, r Request) {
	// todo
}
