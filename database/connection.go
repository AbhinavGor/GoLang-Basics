package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:@/jwt_auth"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database!")
	}
}
