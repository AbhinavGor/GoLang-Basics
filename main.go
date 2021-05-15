package main

import (
	"github.com/gofiber/fiber/v2"
	"com.abhinavgor.test/database"
	"com.abhinavgor.test/routes"
)

func main() {
	database.Connect();

	app := fiber.New()

	routes.Setup(app)

	err := app.Listen(":3000")

	if err != nil{
		panic(err)
	}
}
