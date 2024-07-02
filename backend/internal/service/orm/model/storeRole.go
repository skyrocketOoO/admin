package model

import "gorm.io/gorm"

type (
	Role struct {
		gorm.Model
		Name string `gorm:"unique"`
		Entry
		Permissions []*Permission `gorm:"many2many:RolePermissions;"`
	}
)
