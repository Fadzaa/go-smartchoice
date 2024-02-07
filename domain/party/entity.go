package party

import (
	"gorm.io/gorm"
)

type Party struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"maxLength:100"`
	Image       string `json:"image"`
	Description string `json:"description"`
}
