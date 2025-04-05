package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	CategoryDocumentId uint `gorm:"index"`
	CategoryDocument    CategoryDocument
	ContentDocument    []ContentDocument
	FileStorage        []FileStorage `gorm:"many2many:document_file_storages"`
	Name               string
	ContentRaw         string
}
