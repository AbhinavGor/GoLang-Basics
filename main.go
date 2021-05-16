package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"com.abhinavgor.test/database"
	"com.abhinavgor.test/routes"
)

func main() {
	database.Connect();

	app := fiber.New()
	app.Use(cors.New(cors.Config{
			AllowCredentials: true,
		}))

	routes.Setup(app)

	err := app.Listen(":3000")

	if err != nil{
		panic(err)
	}
}
