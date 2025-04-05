package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type CategoryDocument struct {
	gorm.Model
	Document []Document
	Name string
	Icon sql.NullString
	IsPrivate bool
}
