package party

import (
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	Name        string `json:"name" gorm:"maxLength:100"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
