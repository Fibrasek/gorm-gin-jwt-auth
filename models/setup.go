package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {

	err := godotenv.Load(".env")

	dsn := "host=localhost user=postgres password=postgres dbname=bookstore port=5432 sslmode=disable"
	database, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database! %v", err)
	}

	database.AutoMigrate(&Book{}, &User{})

	DB = database
}
