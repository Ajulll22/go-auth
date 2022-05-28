package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-auth/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-auth?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DB connection error")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
