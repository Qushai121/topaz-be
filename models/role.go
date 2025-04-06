package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	User        []User `gorm:"many2many:role_users"`
	Name        string `gorm:"notNull;size:50"`
	IsDefault   bool   `gorm:"notNull;default:false"`
	Description string `gorm:"notNull"`
}
