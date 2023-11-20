package route

import (
	"atehafinancetracker/controller"

	"github.com/gofiber/fiber/v2"
)



func RouteInit(r *fiber.App)  {
	r.Get("/", func(ctx *fiber.Ctx) error {

		return ctx.SendString("Hello, World 👋!")
})

r.Post("/register", controller.UserControllerRegister)


}