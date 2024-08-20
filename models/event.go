package models

import (
	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Title       string `gorm:"size:100;not null"`
	Description string
	Date        string `gorm:"not null"`
	UserID      uint
}
