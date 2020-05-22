package blog

func NewService(templatesDir string) *Service {
	s := &Service{
		DB:           NewBlog(),
		TemplatesDir: templatesDir,
	}
	return s
}

type Service struct {
	DB           Blog
	TemplatesDir string
}
