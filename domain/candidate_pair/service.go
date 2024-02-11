package candidate_pair

type Service interface {
	//GetCandidatePairByID(id int) (CandidatePair, party.Party, error)
	GetCandidatePairByID(id int) (CandidatePair, error)
}

type ServiceImpl struct {
	candidatePairRepository Repository
}

func NewCandidatePairService(candidatePairRepository Repository) *ServiceImpl {
	return &ServiceImpl{candidatePairRepository}
}

//func (s *ServiceImpl) GetCandidatePairByID(id int) (CandidatePair, party.Party, error) {
//	candidatePair, party, err := s.candidatePairRepository.FindByID(id)
//	if err != nil {
//		return candidatePair, party, err
//	}
//
//	return candidatePair, party, nil
//}

func (s *ServiceImpl) GetCandidatePairByID(id int) (CandidatePair, error) {
	candidatePair, err := s.candidatePairRepository.FindByID(id)
	if err != nil {
		return candidatePair, err
	}

	return candidatePair, nil
}
