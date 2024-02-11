package debate

type Debate struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"maxLength:100"`
	Description string `json:"description" gorm:"maxLength:100"`
	Date        string `json:"date" gorm:"maxLength:100"`
	Link        string `json:"link"`
}
