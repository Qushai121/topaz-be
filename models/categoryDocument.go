package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type CategoryDocument struct {
	gorm.Model
	Document  []Document
	Name      string         `gorm:"notNull;size:100"`
	Icon      sql.NullString `gorm:"notNull;default:default"`
	IsPrivate bool           `gorm:"notNull;default:false"`
}
