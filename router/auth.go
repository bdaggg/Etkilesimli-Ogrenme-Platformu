package router

import (
	"EOP/controllers"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")

	user.Post("/", controllers.Singup)
	user.Post("/login", controllers.Login)
	user.Post("/update", controllers.Update)

}
