package blog

func NewService(templatesDir string) *Service {
	s := &Service{
		DB:           NewPage(),
		TemplatesDir: templatesDir,
	}
	return s
}

type Service struct {
	DB           Page
	TemplatesDir string
}
