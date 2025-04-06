package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Title   string `gorm:"notNull;size:55"`
	Summary string `gorm:"notNull;size:255"`
	NewsId  uint   `gorm:"index"`
	News    News
}
