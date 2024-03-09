package controllers

import (
	"EOP/database"
	"EOP/helpers"
	"EOP/middleware"
	"EOP/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func SingupTeach(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.Personel)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error was able to body parse try again ", "data": err})
	}
	// username controll
	err = helpers.Username_controll(user.UserName)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "username used by another one", "data": err})
	}

	// mail controll
	err = helpers.Mail_Control(user.Mail)
	log.Println(err)
	if err == nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "mail used by another one", "data": err})
	}

	user.Password = helpers.HashPass(user.Password)
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error was able to create new user try again", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user created", "data": user})

}

func LoginTeach(c *fiber.Ctx) error {
	logindata := new(model.Login)
	userdata := new(model.Personel)
	db := database.DB.Db
	c.BodyParser(&logindata)
	logindata.Password = helpers.HashPass(logindata.Password)
	err := db.Where("user_name = ? and password = ?", logindata.UserName, logindata.Password).First(&userdata).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "faild", "message": "faild login"})
	}

	token, err := middleware.CreateToken(userdata.UserName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token oluşturulamadı")
	}

	session := new(model.Session)
	session.UserID = userdata.UserID
	session.Token = token
	db.Save(&session)
	helpers.LogMessage(userdata.UserName, "- giriş yapti")
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login success", "data": userdata, "token": token})

}

func UpdatePasswordTeach(c *fiber.Ctx) error {
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
	ChangePassword.OldPassword = helpers.HashPass(ChangePassword.OldPassword)
	if user.Password != ChangePassword.OldPassword {
		return c.Status(400).JSON(fiber.Map{"status": "faild", "message": "old password is false"})
	}
	user.Password = helpers.HashPass(ChangePassword.NewPassword)

	err = database.DB.Db.Where("user_name=?", user.UserName).Updates(&user).Error
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "faild", "message": "faild login"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "password update success", "data": user})
}

func Logout(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	userID := user.UserID
	var userSession []model.Session
	err = database.DB.Db.Where("user_id =?", userID).Find(&userSession).Error
	if err != nil {
		return err
	}
	for _, session := range userSession {
		err = database.DB.Db.Delete(&session).Error
		if err != nil {
			return c.Status(402).JSON(fiber.Map{"status": "error", "message": "logout error"})
		}
	}
	/*name := c.Locals("user").(model.Personel).UserName
	log.Println(name) */
	//yukarıdaki işlem log kaydı için name bilgisini tutar

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "logout success"})
}

//todo apidog yukarıdaki son iki enpoint için api req oluştur
