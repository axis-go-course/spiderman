package blog

func NewAuthor(name string) *Author {
	return &Author{name: name}
}

type Author struct {
	name string
}

func (a *Author) WriteArticle(page Page, article *Article) error {
	return page.saveArticle(article)
}

func (a *Author) EraseArticle(page Page, article *Article) error {
	return page.DeleteArticle(article.Title)
}
