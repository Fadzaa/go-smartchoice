package debate

import "gorm.io/gorm"

type Repository interface {
	GetAllDebate() ([]Debate, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewDebateRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) GetAllDebate() ([]Debate, error) {
	var debates []Debate

	err := r.db.Find(&debates).Error

	if err != nil {
		return nil, err
	}

	return debates, nil
}
