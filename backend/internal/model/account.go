package model

import (
	"gorm.io/gorm"
)

type (
	Account struct {
		gorm.Model
		UserName    string `gorm:"unique;not null"`
		State       int32  `gorm:"not null;default:1"` // 1: active, 2:inactive
		HashPass    []byte `gorm:"not null"`
		Salt        []byte `gorm:"not null"`
		DisplayName string `gorm:"not null"`
		// Phone       string `gorm:"unique"`
		// Email string `gorm:"unique"`

		RoleID uint
		Role   Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)
