package blog

import "testing"

func Test_author_writes_article(t *testing.T) {
	author := &Author{name: "Peter Parker"}
	page := NewPage()
	article := &Article{Title: "Helping Iron Man", Content: "wip"}
	if err := author.WriteArticle(page, article); err != nil {
		t.Error(err)
	}
}
