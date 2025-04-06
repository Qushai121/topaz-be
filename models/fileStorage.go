package models

import "gorm.io/gorm"

type FileStorage struct {
	gorm.Model
	Filename string `gorm:"notNull"`
	FileExt  string `gorm:"notNull"`
	FilePath string `gorm:"notNull"`
}
