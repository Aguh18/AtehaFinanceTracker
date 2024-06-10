package main

import (
	"atehafinancetracker/database"
	"atehafinancetracker/database/migration"
	_ "atehafinancetracker/docs"
	"atehafinancetracker/route"

	"github.com/gofiber/fiber/v2"
)

//	@title			Finance Tracker API
//	@description	This API is for tracking your finance
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.email	fiber@swagger.io
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@BasePath		/
// 	@host			localhost:8080
func main()  {

	app := fiber.New()

	// db connection
	database.DatabaseInit()
	// migarate
	migration.Migrate()

	route.RouteInit(app)

	
	

	app.Listen(":8080")
	
}