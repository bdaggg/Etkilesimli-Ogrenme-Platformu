package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentID int    `gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name,omitempty"`
	Surname   string `json:"surname,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	Password  string `json:"password,omitempty"`
	Mail      string `json:"mail,omitempty"`
	Job       string `json:"job,omitempty"`
}

type StudentLogin struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type StudentUpdate struct {
	Token string `json:"token"`
}
