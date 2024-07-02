package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Account struct {
		gorm.Model
		Username string `gorm:"unique;not null"`
		HashPass string `gorm:"not null"`
		Entry
	}

	Entry struct {
		EntryID   uuid.UUID `gorm:"not null"`
		EntryType EntryType `gorm:"not null"`
	}
)

type EntryType string

const (
	StoreEntry EntryType = "Store"
	FirmEntry  EntryType = "Firm"
)
