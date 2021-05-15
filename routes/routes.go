package routes

import (
	"github.com/gofiber/fiber/v2"
	"com.abhinavgor.test/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
}
