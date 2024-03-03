package controllers

import (
	"EOP/database"
	"EOP/helpers"
	"EOP/middleware"
	"EOP/model"

	"github.com/gofiber/fiber/v2"
)

func LoginStudemt(c *fiber.Ctx) error {
	logindata := new(model.StudentLogin)
	userdata := new(model.Student)
	db := database.DB.Db
	c.BodyParser(&logindata)
	err := db.Where("user_name = ? and password = ?", logindata.UserName, logindata.Password).First(&userdata).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "faild", "message": "faild login"})
	}

	token, err := middleware.CreateToken(userdata.UserName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token oluşturulamadı")
	}

	session := new(model.Session)
	session.UserID = userdata.StudentID
	session.Token = token
	db.Save(&session)
	helpers.LogMessage(userdata.UserName, "- giriş yapti")
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login success", "data": userdata, "token": token})

}

func UpdatePasswordStudent(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	type changePassword struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_Password"`
	}
	var ChangePassword changePassword
	c.BodyParser(&ChangePassword)
	if user.Password != ChangePassword.OldPassword {
		return c.Status(400).JSON(fiber.Map{"status": "faild", "message": "old password is false"})
	}
	user.Password = ChangePassword.NewPassword

	err = database.DB.Db.Where("user_name=?", user.UserName).Update("password", user.Password).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "faild", "message": "faild login"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "password update success", "data": user})
}
