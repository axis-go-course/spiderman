package blog

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

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
