package controllers

import (
	"github.com/gofiber/fiber/v2"
	"com.abhinavgor.test/models"
	"golang.org/x/crypto/bcrypt"
	"com.abhinavgor.test/database"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
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
			"message": "User not found!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid credentials!",
		})
	}

	claims := jwt.NewWithClaims(jwt.SignInMethodSH256, jwt.StandardClaims{
		Issuer: strconv.Itoa(user.Id),
		ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
	})

	token, err := claims.SignedString([]byte("SecretKey"))

	if err != nil{
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login!",
		})
	}

	return c.JSON(user)
}
