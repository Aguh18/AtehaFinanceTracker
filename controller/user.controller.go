package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func UserControllerRegister(ctx *fiber.Ctx) error {

	registerUser := new(request.User)

	if err := ctx.BodyParser(registerUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	valdate := validator.New()

	if err := valdate.Struct(registerUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation error",
			"error":   err.Error(),
		})
	}
	NewUser := entity.User{
		Username: registerUser.Username,
		Name:    registerUser.Name,
		Email:   registerUser.Email,
		Password: registerUser.Password,

	}

	if err := database.DB.Create(&NewUser).Error; err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success create User",
		"data":    NewUser,
	})
}
