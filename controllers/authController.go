package controllers

import (
	"github.com/gofiber/fiber/v2"
	"com.abhinavgor.test/models"
	"golang.org/x/crypto/bcrypt"
	"com.abhinavgor.test/database"
)

func Register (c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login (c *fiber.Ctx) error {
	var data map[string]string

	if err :=  c.BodyParser(&data); err != nil {
		panic(err)
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0{
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not fond!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials!",
		})
	}

	return c.JSON(user)
}

