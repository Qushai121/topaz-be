package models

import "gorm.io/gorm"

type FileStorage struct {
	gorm.Model
	Filename string
	FileExt string
	FilePath string
}