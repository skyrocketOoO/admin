package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Role struct {
		gorm.Model
		StoreID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_store_name"`
		Name    string    `gorm:"not null;uniqueIndex:idx_store_name"`
		Page    Page      `gorm:"embedded; embeddedPrefix:Page"`
	}

	PageOperation struct {
		Create bool `gorm:"default:false"`
		Read   bool `gorm:"default:false"`
		Update bool `gorm:"default:false"`
		Delete bool `gorm:"default:false"`
	}

	Page struct {
		LoginRecord PageOperation `gorm:"embedded;embeddedPrefix:LoginRecord"`
		Role        PageOperation `gorm:"embedded;embeddedPrefix:Role"`
		Setting     PageOperation `gorm:"embedded;embeddedPrefix:Setting"`
		AccountList PageOperation `gorm:"embedded;embeddedPrefix:AccountList"`
	}
)
