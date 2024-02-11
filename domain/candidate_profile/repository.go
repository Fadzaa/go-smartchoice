package candidate_profile

import "gorm.io/gorm"

type Repository interface {
	FindByID(id int) (CandidateProfile, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewCandidateProfileRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) FindByID(id int) (CandidateProfile, error) {
	var candidateProfile CandidateProfile
	err := r.db.Preload("Achievements").First(&candidateProfile, id).Error
	return candidateProfile, err
}
