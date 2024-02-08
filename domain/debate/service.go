package debate

type Service interface {
	GetAllDebate() ([]Debate, error)
}

type ServiceImpl struct {
	repositoryDebate Repository
}

func NewDebateService(repositoryDebate Repository) *ServiceImpl {
	return &ServiceImpl{repositoryDebate}
}

func (s *ServiceImpl) GetAllDebate() ([]Debate, error) {
	return s.repositoryDebate.GetAllDebate()
}
