package controllers

import (
	"EOP/database"
	"EOP/model"

	"github.com/gofiber/fiber/v2"
)

func SessionController(c *fiber.Ctx, token string) (model.User, error) {
	db := database.DB.Db
	sesion := new(model.Session)
	sesion.Token = token
	err := db.Where("token=?", sesion.Token).First(&sesion).Error
	if err != nil {
		return model.User{}, c.Status(400).JSON(fiber.Map{"status": "faild", "message": "session not found"})
	}
	user := new(model.User)
	user.UserID = sesion.UserID
	err = db.Where("user_id=?", user.UserID).First(&user).Error
	if err != nil {
		return model.User{}, c.Status(400).JSON(fiber.Map{"status": "faild", "message": "user not found"})
	}

	return *user, nil
}
