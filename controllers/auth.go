package controllers

import (
	"EOP/database"
	"EOP/helpers"
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
	//log.Println(reflect.TypeOf(user.Name))
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
