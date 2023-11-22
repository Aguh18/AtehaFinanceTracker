package route

import (
	"atehafinancetracker/controller"
	"atehafinancetracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	// User Routes
	r.Post("/register", controller.UserControllerRegister)

	// Acount Routes
	r.Post("/acount/create", middleware.Middelware ,controller.AcountControllerCreate)

	// Auth Routes
	r.Post("/login", controller.Login)
}
