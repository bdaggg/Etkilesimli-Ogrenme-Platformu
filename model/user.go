package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Role     string `json:"role"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Job      string `json:"job"`
}
