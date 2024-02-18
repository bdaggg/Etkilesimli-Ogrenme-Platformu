package main

import (
	"EOP/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Listen(":3000")

}
