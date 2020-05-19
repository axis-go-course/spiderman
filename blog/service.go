package blog

func NewService(templatesDir string) *Service {
	s := &Service{
		blog:         NewBlog(),
		templatesDir: templatesDir,
	}
	return s
}

type Service struct {
	blog         Blog
	templatesDir string
}
