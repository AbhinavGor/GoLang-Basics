package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"com.abhinavgor.test/models"
)

var DB *gorm.DB

func Connect() {
	conn, err := gorm.Open(mysql.Open("root:@/jwt_auth"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}
	
	DB = conn

	conn.AutoMigrate(&models.User{})
}
