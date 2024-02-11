package party

type Party struct {
	ID              uint   `json:"id" gorm:"primaryKey"`
	Name            string `json:"name" gorm:"maxLength:100"`
	Image           string `json:"image"`
	Description     string `json:"description"`
	CandidatePairID uint   `json:"candidate_pair_id"`
}
