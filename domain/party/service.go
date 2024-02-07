package party

type Service interface {
	GetAll() ([]Party, error)
	Search(query string) ([]Party, error)
}

type ServiceImpl struct {
	repositoryParty Repository
}

func NewPartyService(repositoryParty Repository) *ServiceImpl {
	return &ServiceImpl{repositoryParty}
}

func (s *ServiceImpl) GetAll() ([]Party, error) {
	return s.repositoryParty.FindAll()
}

func (s *ServiceImpl) Search(query string) ([]Party, error) {
	return s.repositoryParty.Search(query)
}
