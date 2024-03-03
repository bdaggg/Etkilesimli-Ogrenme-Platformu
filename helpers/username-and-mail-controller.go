package helpers

import (
	"EOP/database"
	"EOP/model"
)

func Username_controll(user_name string) error {
	db := database.DB.Db
	//username control
	var existingUser model.Personel
	result := db.Where("user_name = ?", user_name).First(&existingUser).Error

	if result != nil {
		return result
	}
	return nil
}

func Mail_Control(mail string) error {
	db := database.DB.Db
	var existingUser model.Personel
	//email control
	result := db.Where("mail = ?", mail).First(&existingUser).Error

	if result != nil {
		return result
	}
	return nil
}
func Student_username_controll(user_name string) error {
	db := database.DB.Db
	//username control
	var existingUser model.Personel
	result := db.Where("user_name = ?", user_name).First(&existingUser).Error

	if result != nil {
		return result
	}
	return nil
}

func Student_mail_Control(mail string) error {
	db := database.DB.Db
	var existingUser model.Personel
	//email control
	result := db.Where("mail = ?", mail).First(&existingUser).Error

	if result != nil {
		return result
	}
	return nil
}
