package blog

import "fmt"

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewBlog() Blog {
	return make(map[string]*Article)
}

type Blog map[string]*Article

func (b Blog) SaveArticle(v ...*Article) error {
	for _, a := range v {
		if a.Title == "" {
			return fmt.Errorf("missing title")
		}
		b[a.Title] = a
	}
	return nil
}

func (b Blog) LoadArticles(v []*Article) int {
	var i int
	for _, a := range b {
		if i == len(v) {
			break
		}
		v[i] = a
		i++
	}
	return i
}

func (b Blog) DeleteArticle(title string) int {
	if _, found := b[title]; found {
		delete(b, title)
		return 1
	}
	return 0
}
