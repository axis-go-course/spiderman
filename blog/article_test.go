package blog

import "testing"

func TestSaveArticle(t *testing.T) {
	b := NewBlog()
	a := &Article{Title: "not empty", Content: "not empty"}
	if err := b.SaveArticle(a); err != nil {
		t.Error(err)
	}
}

func TestLoadArticles(t *testing.T) {
	b := NewBlog()
	for _, a := range []*Article{
		&Article{Title: "first", Content: "first"},
		&Article{Title: "second", Content: "second"},
	} {
		b.SaveArticle(a)
	}

	result := make([]*Article, 3)
	if got, exp := b.LoadArticles(result), 2; got != exp {
		t.Error("expected", exp, "articles, got", got)
	}

	result = make([]*Article, 1)
	if got, exp := b.LoadArticles(result), 1; got != exp {
		t.Error("expected", exp, "articles, got", got)
	}
}

func TestDeleteArticle(t *testing.T) {
	b := NewBlog()
	a := &Article{Title: "first", Content: "first"}
	b.SaveArticle(a)
	if got, exp := b.DeleteArticle(a.Title), 1; got != exp {
		t.Error("expected", exp, "articles, got", got)
	}

	if got, exp := b.DeleteArticle("no such title"), 0; got != exp {
		t.Error("expected", exp, "articles, got", got)
	}
}
