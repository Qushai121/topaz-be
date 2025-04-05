package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Title   string
	Summary string
	NewsId  uint `gorm:"index"`
	News News
}
