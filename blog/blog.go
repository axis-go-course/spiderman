package blog

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/pborman/uuid"
)

func NewDatabase() Database {
	return make(map[string]*Article)
}

type Database map[string]*Article

type Article struct {
	Id      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
}

func (a *Article) Reader() io.Reader {
	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(a)
		w.Close()
	}()
	return r
}

func (b Database) SaveArticle(v ...*Article) error {
	for _, a := range v {
		if a.Title == "" {
			return fmt.Errorf("missing title")
		}
		a.Id = uuid.NewUUID()
		b[a.Title] = a
	}
	return nil
}

func (b Database) LoadArticles(v []*Article) int {
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

func (b Database) DeleteArticle(title string) error {
	if _, found := b[title]; found {
		delete(b, title)
		return nil
	}
	return fmt.Errorf("article not found")
}
