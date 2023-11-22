package middleware

import (
	"atehafinancetracker/utils"

	"github.com/gofiber/fiber/v2"
)

func Middelware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")

	if token == "" {
		return ctx.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized jwt emty",
		})
	}

	claims, err := utils.DecodeToken(token)
	if err != nil{
		return ctx.Status(401).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	ctx.Locals("user", claims)

	return ctx.Next()

}
