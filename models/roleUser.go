package models

import (
	"time"

	"gorm.io/gorm"
)

type RoleUser struct {
	RoleId    uint `gorm:"primarykey"`
	UserId    uint `gorm:"primarykey"`
	View      bool `gorm:"notNull;default:false"`
	Create    bool `gorm:"notNull;default:false"`
	Update    bool `gorm:"notNull;default:false"`
	Delete    bool `gorm:"notNull;default:false"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
