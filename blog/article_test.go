package blog

import "testing"

func Test_blog(t *testing.T) {
	A := &Article{Title: "first", Content: "first"}
	B := &Article{Title: "second", Content: "second"}
	all := []*Article{A, B}
	b := NewBlog()
	b.SaveArticle(all...)

	b.SaveArticle(&Article{})        // should never get saved
	b.DeleteArticle("no such title") // nothing removed

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
