package main

import (
	"./database"
	"./routes"
)

func main() {
	database.connection()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000");
}
