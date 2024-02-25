package controllers

import (
	"EOP/database"
	"EOP/helpers"
	"EOP/middleware"
	"EOP/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Singup(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)
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

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "server error was able to create new user try again", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "user created", "data": user})

}

func Login(c *fiber.Ctx) error {
	logindata := new(model.Login)
	userdata := new(model.User)
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
	session.UserID = userdata.UserID
	session.Token = token
	db.Save(&session)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "login success", "data": userdata, "token": token})

}

func Update(c *fiber.Ctx) error {
	user, err := middleware.TokenControl(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "data": err})
	}
	log.Println(user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "procces success", "data": user})
}
