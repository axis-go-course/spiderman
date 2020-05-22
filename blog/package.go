/*
Package blog provides domain abstractions for writing blog articles.

Example

	author := &Author{name: "Peter Parker"}
	page := NewPage()
	article := &Article{Title: "Helping Iron Man", Content: "wip"}
	err := author.WriteArticle(page, article)

*/
package blog
