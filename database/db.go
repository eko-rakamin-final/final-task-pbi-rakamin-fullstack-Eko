package database

import (
	"log"
	"rakamin-go/app"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/rakamin_api"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	
	database.AutoMigrate(&app.User{}, &app.Photo{})
	DB = database
	log.Println("Database connected!")
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

