package models

import (
	"time"

	"gorm.io/gorm"
)

type DocumentFileStorage struct {
	DocumentId    uint `gorm:"primarykey"`
	FileStorageId uint `gorm:"primarykey"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
