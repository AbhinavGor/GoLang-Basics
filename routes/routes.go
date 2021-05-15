package routes

import (
	"github.com/gofiber/fiber/v2"
	"com.abhinavgor.test/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)

	app.Post("/api/Login", controllers.Login)
}
