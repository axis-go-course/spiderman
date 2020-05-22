package blog

import (
	"testing"
)

func Test_blog(t *testing.T) {
	A := &Article{Title: "first", Content: "first"}
	B := &Article{Title: "second", Content: "second"}
	all := []*Article{A, B}
	page := NewPage()
	author := NewAuthor("Peter Parker")
	page.saveArticle(all...)
	mustNot(t, author.WriteArticle(page, &Article{}))
	mustNot(t, page.DeleteArticle("no such title"))

	for _, v := range all {
		if v.Id.String() == "" {
			t.Fatal("UUID not generated")
		}
	}

	result := make([]*Article, 5)
	exp := len(all)
	if got := page.LoadArticles(result); got != exp {
		t.Error("load all expected", exp, "articles, got", got)
	}

	exp = 1
	if got := page.LoadArticles(result[:exp]); got != exp {
		t.Error("load partial expected", exp, "articles, got", got)
	}

	if err := author.EraseArticle(page, A); err != nil {
		t.Error("delete expected", err)
	}
}

func mustNot(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Fatal(err)
	}
}
