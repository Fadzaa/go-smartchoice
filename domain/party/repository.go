package party

import "gorm.io/gorm"

type Repository interface {
	GetAllParty(query string) ([]Party, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewPartyRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

func (r *RepositoryImpl) GetAllParty(query string) ([]Party, error) {
	var parties []Party
	err := r.db.Where("name LIKE ?", "%"+query+"%").Find(&parties).Error
	if err != nil {
		return nil, err
	}
	return parties, nil
}
