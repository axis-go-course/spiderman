package blog

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_blog(t *testing.T) {
	A := &Article{Title: "first", Content: "first"}
	B := &Article{Title: "second", Content: "second"}
	all := []*Article{A, B}
	b := NewBlog()
	b.SaveArticle(all...)
	mustNot(t, b.SaveArticle(&Article{}))
	mustNot(t, b.DeleteArticle("no such title"))

	for _, v := range all {
		if v.Id.String() == "" {
			t.Fatal("UUID not generated")
		}
	}

	result := make([]*Article, 5)
	exp := len(all)
	if got := b.LoadArticles(result); got != exp {
		t.Error("load all expected", exp, "articles, got", got)
	}

	exp = 1
	if got := b.LoadArticles(result[:exp]); got != exp {
		t.Error("load partial expected", exp, "articles, got", got)
	}

	if err := b.DeleteArticle(A.Title); err != nil {
		t.Error("delete expected", err)
	}
}

func Test_article(t *testing.T) {
	A := &Article{Title: "first", Content: "first"}
	var buf bytes.Buffer
	io.Copy(&buf, A.Reader())
	got := buf.String()
	if !strings.Contains(got, "title\":") {
		t.Fatal("missing title", got)
	}
}

func mustNot(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal(err)
	}
}
