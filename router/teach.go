package router

import (
	"EOP/controllers"

	"github.com/gofiber/fiber/v2"
)

func Teach(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	teach := v1.Group("/taech")

	teach.Post("/", controllers.AddQuestions)
}
