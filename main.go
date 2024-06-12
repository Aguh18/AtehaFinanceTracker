package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"atehafinancetracker/database"
	"atehafinancetracker/database/migration"
	_ "atehafinancetracker/docs"
	"atehafinancetracker/route"
)

//	@title			Finance Tracker API
//	@description	This API is for tracking your finance
//	@contact.name	API Support
//	@contact.email	teguh180902@gmail.com
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@BasePath		/
// 	@host			localhost:8080
func main()  {

	app := fiber.New()

	app.Use(cors.New())

	// db connection
	database.DatabaseInit()
	// migarate
	migration.Migrate()

	route.RouteInit(app)

	
	

	app.Listen(":8080")
	
}