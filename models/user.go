package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role []Role `gorm:"many2many:role_users"`
	News []News
	FirstName string
	LastName sql.NullString
	Email string
	Password string
	RememberToken string 
}
