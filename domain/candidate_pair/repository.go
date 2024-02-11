package candidate_pair

import (
	"gorm.io/gorm"
)

type Repository interface {
	//FindByID(id int) (CandidatePair, party.Party, error)
	FindByID(id int) (CandidatePair, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewCandidatePairRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) FindByID(id int) (CandidatePair, error) {
	var candidatePair CandidatePair
	err := r.db.Preload("Parties").First(&candidatePair, id).Error

	if err != nil {
		return candidatePair, err
	}

	return candidatePair, nil
}
