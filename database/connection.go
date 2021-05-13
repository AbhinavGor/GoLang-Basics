package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connect() {
	_, err := gorm.Open(mysql.Open("root:@/jwt_auth"), &gorm.Config{})

	if err != nil{
		panic("Could not connect to database!")
	}

}
