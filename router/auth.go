package router

import (
	"EOP/controllers"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user")

	user.Get("/", controllers.Logout)
	user.Post("/", controllers.SingupTeach)
	user.Post("/login", controllers.LoginTeach)
	user.Put("/", controllers.UpdatePasswordTeach)

}
