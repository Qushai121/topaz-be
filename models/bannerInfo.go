package models

import "gorm.io/gorm"

type BannerInfo struct {
	gorm.Model
	NewsId string `gorm:"index"`
	News News
	Title string
	Link string
	Icon string
	PriorityScale string
	Severity string
}
