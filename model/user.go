package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   uint   `gorm:"primaryKey;autoIncrement`
	Role     string `json:"role,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
	Mail     string `json:"mail,omitempty"`
	Job      string `json:"job,omitempty"`
}

type Login struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}
