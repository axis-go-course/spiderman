package blog

func NewService(templatesDir string) *Service {
	s := &Service{
		Database:     NewDatabase(),
		TemplatesDir: templatesDir,
	}
	return s
}

type Service struct {
	Database     Database
	TemplatesDir string
}
