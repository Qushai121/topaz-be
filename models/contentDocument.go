package models

import "gorm.io/gorm"

type ContentDocument struct {
	gorm.Model
	DocumentId uint `gorm:"index"`
	Name string
	Body string
}