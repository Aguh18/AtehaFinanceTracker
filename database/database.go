package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func DatabaseInit() {
	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}
	var err error
	PostgressCon := os.Getenv("DatabaseCon")
	dsn := PostgressCon
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")
	}
	log.Println("Connection Opened to Database")
}
