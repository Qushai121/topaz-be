package models

import (
	"time"

	"gorm.io/gorm"
)

type RoleUser struct {
	RoleId    uint `gorm:"primarykey"`
	UserId    uint `gorm:"primarykey"`
	View      bool
	Create    bool
	Update    bool
	Delete    bool
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
