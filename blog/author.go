package blog

type Author struct {
	name string
}

func (a *Author) WriteArticle(page Page, article *Article) error {
	return page.SaveArticle(article)
}
