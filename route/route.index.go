package route

import (
	"atehafinancetracker/controller"

	"github.com/gofiber/fiber/v2"
)



func RouteInit(r *fiber.App)  {
	
// User Routes
r.Post("/register", controller.UserControllerRegister)




// Acount Routes
r.Post("/acount/create", controller.AcountControllerCreate)


}