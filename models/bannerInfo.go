package models

import "gorm.io/gorm"

type BannerInfo struct {
	gorm.Model
	NewsId        string `gorm:"index"`
	News          News
	Title         string `gorm:"notNull;size:100"`
	Link          string `gorm:"notNull"`
	Icon          string `gorm:"notNull;default:default"`
	PriorityScale string `gorm:"notNull"`
	Severity      string `gorm:"notNull;default:info"`
}
