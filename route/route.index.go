package route

import (
	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"

	"atehafinancetracker/controller"
	_ "atehafinancetracker/docs"
	"atehafinancetracker/middleware"
)



func RouteInit(r *fiber.App) {



	// @Tags Register
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
	// transaction create
	r.Post("/transaction/create", middleware.Middelware, controller.TransactionControllerCreate)
	// get all transaction by acount id
	r.Get("/transaction/:id", middleware.Middelware ,controller.TransactionControllerGetByAcountId)


	r.Get("/swagger/*", swagger.HandlerDefault) // default

	r.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL: "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))


	

}
