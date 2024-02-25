package middleware

import (
	"EOP/model"
	"log"

	"github.com/gofiber/fiber/v2"
)

func TokenControl(c *fiber.Ctx) (model.User, error) {
	authorizationHeader := c.Get("Authorization")

	if authorizationHeader == "" || len(authorizationHeader) < 7 || authorizationHeader[:7] != "Bearer " {
		return model.User{}, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or missing token",
		})
	}
	token := authorizationHeader[7:]
	log.Println(token)
	//todo gelen tokeni databasenden control et ve useri dön
	return model.User{}, nil

}
