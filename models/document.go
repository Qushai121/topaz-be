package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	CategoryDocumentId uint `gorm:"index"`
	CategoryDocument   CategoryDocument
	ContentDocument    []ContentDocument
	FileStorage        []FileStorage `gorm:"many2many:document_file_storages"`
	Name               string        `gorm:"notNull;size:100"`
	ContentRaw         string        `gorm:"notNull"`
}
