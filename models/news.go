package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	UserId uint `gorm:"index"`
	NotificationId uint `gorm:"index"`
	Title string
	Body string
	IsSendNotification bool
	IsPublish bool
}
