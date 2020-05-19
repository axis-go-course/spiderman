package blog

func NewService(templatesDir string) *Service {
	s := &Service{
		Blog:         NewBlog(),
		TemplatesDir: templatesDir,
	}
	return s
}

type Service struct {
	Blog         Blog
	TemplatesDir string
}
