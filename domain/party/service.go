package party

type Service interface {
	GetAllParty(query string) ([]Party, error)
}

type ServiceImpl struct {
	repositoryParty Repository
}

func NewPartyService(repositoryParty Repository) *ServiceImpl {
	return &ServiceImpl{repositoryParty}
}

func (s *ServiceImpl) GetAllParty(query string) ([]Party, error) {
	return s.repositoryParty.GetAllParty(query)
}
