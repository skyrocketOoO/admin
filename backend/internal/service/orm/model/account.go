package model

import (
	"gorm.io/gorm"
)

type (
	Account struct {
		gorm.Model
		UserName    string `gorm:"unique;not null"`
		HashPass    []byte `gorm:"not null"`
		Salt        string `gorm:"size:255;not null"`
		DisplayName string `gorm:"not null"`
		// Phone       string `gorm:"unique"`
		Email string `gorm:"unique"`
	}
)
