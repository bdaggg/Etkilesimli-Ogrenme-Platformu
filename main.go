package main

import (
	"EOP/database"
	"EOP/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	router.Auth(app)
	router.Teach(app)

	app.Listen(":3000")
}
