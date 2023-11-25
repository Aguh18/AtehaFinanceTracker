package route

import (
	"atehafinancetracker/controller"
	"atehafinancetracker/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	// User Routes
	r.Post("/register", controller.UserControllerRegister)
	// get all user
	r.Get("/user",  controller.UserControllerGetAll)
	// update user
	r.Put("/user/:id", middleware.Middelware ,controller.UserControllereditUser)

	// Acount Routes

	// create a new account
	r.Post("/acount/create", middleware.Middelware ,controller.AcountControllerCreate)
	// get all accounts byid
	r.Get("/acount/", middleware.Middelware ,controller.AcountControllerGetByUserId)
	// update data acount
	r.Put("/acount/:id", middleware.Middelware ,controller.AcountControllerUpdate)
	// delete acount
	r.Delete("/acount/:id", middleware.Middelware ,controller.AcountControllerDeleteById)


	// Auth Routes
	r.Post("/login", controller.Login)


	// Transaction Routes
	r.Post("/transaction/create", controller.TransactionControllerCreate)
}
