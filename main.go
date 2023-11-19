package main

import (
	"atehafinancetracker/database"
	"atehafinancetracker/database/migration"
	"atehafinancetracker/route"

	"github.com/gofiber/fiber/v2"
)


func main()  {

	app := fiber.New()

	// db connection
	database.DatabaseInit()
	// migarate
	migration.Migrate()



	// init route
	route.RouteInit(app)


	app.Listen(":8080")
	
}