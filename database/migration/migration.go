package migration

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"log"
)


func Migrate()  {

	err := database.DB.AutoMigrate(&entity.User{}, entity.Acount{}, entity.Transaction{})
	if err != nil {
		log.Fatal("Error migrating database")
	}else{
		log.Println("Database migrated successfully")
	}
	
}