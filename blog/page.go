package blog

import (
	"fmt"

	"github.com/google/uuid"
)

func NewPage() Page {
	return make(map[string]*Article)
}

type Page map[string]*Article

func (p Page) saveArticle(v ...*Article) error {
	for _, a := range v {
		if a.Title == "" {
			return fmt.Errorf("missing title")
		}
		a.Id, _ = uuid.NewUUID()
		p[a.Title] = a
	}
	return nil
}

// LoadArticles loads the first articles into v, returns number of
// articles loaded.
func (p Page) LoadArticles(v []*Article) int {
	var i int
	for _, a := range p {
		if i == len(v) {
			break
		}
		v[i] = a
		i++
	}
	return i
}

func (p Page) deleteArticle(title string) error {
	if _, found := p[title]; found {
		delete(p, title)
		return nil
	}
	return fmt.Errorf("article not found")
}
