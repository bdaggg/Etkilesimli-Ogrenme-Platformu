package model

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID    uint `gorm:"primaryKey;autoIncrement`
	Token     string
	Is_active bool `gorm:"default:true"`
}
