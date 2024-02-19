package helpers

import (
	"EOP/database"
	"EOP/model"
)

func Username_controll(user_name string) error {
	db := database.DB.Db
	//username control
	var existingUser model.User
	result := db.Where("user_name = ?", user_name).First(&existingUser).Error

	if result != nil {
		return result
	}
	return nil
}

func Mail_Control(mail string) error {
	db := database.DB.Db
	var existingUser model.User
	//email control
	result := db.Where("mail = ?", mail).First(&existingUser).Error

	if result != nil {
		return result
	}
	return nil
}
