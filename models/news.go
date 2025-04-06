package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	UserId             uint   `gorm:"index"`
	NotificationId     uint   `gorm:"index"`
	Title              string `gorm:"notNull;size:55"`
	Body               string `gorm:"notNull"`
	IsSendNotification bool   `gorm:"notNull;default:true"`
	IsPublish          bool   `gorm:"notNull;default:true"`
}
