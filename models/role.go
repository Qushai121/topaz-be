package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	User        []User `gorm:"many2many:role_users"`
	Name        string
	IsDefault   bool
	Description string
}
