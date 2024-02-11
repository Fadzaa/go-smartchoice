package candidate_profile

type CandidateProfile struct {
	ID            uint          `json:"id" gorm:"primaryKey"`
	Name          string        `json:"name" gorm:"maxLength:50"`
	Image         string        `json:"image"`
	Position      string        `json:"position" gorm:"maxLength:100"`
	ShortProfile  string        `json:"short_profile"`
	BirthdayPlace string        `json:"birthday_place" gorm:"maxLength:50"`
	CandidateAs   string        `json:"candidate_as" gorm:"maxLength:50"`
	TotalAsset    string        `json:"total_asset" gorm:"maxLength:50"`
	LinkInstagram string        `json:"link_instagram"`
	LinkTwitter   string        `json:"link_twitter"`
	Achievements  []Achievement `json:"achievements"` // 1 to Many
}

type Achievement struct {
	ID                 int    `json:"id" gorm:"primaryKey"`
	Title              string `json:"title" gorm:"maxLength:100"`
	Description        string `json:"description"`
	CandidateProfileID uint   `json:"candidate_profile_id"`
}
