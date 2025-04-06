package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role          []Role `gorm:"many2many:role_users"`
	News          []News
	FirstName     string         `gorm:"notNull;size:50"`
	LastName      sql.NullString `gorm:"size:50"`
	Email         string         `gorm:"notNull;unique"`
	Password      string         `gorm:"notNull"`
	RememberToken string
}
