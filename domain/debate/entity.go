package debate

import (
	"gorm.io/gorm"
)

type Debate struct {
	gorm.Model
	Name        string `json:"name" gorm:"maxLength:100"`
	Image       string `json:"image"`
	Description string `json:"description" gorm:"maxLength:100"`
	Date        string `json:"date" gorm:"maxLength:100"`
	Link        string `json:"link"`
}
