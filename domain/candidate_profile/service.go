package candidate_profile

type Service interface {
	FindByID(id int) (CandidateProfile, error)
}

type ServiceImpl struct {
	repo Repository
}

func NewCandidateProfileService(repo Repository) *ServiceImpl {
	return &ServiceImpl{repo}
}

func (s *ServiceImpl) FindByID(id int) (CandidateProfile, error) {
	return s.repo.FindByID(id)
}
