package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/utils"

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
	if hashedpass, err := utils.HashPassword(registerUser.Password); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot hash password",
			"error":   err.Error(),
		})
	} else {
		registerUser.Password = hashedpass
	}

	NewUser := entity.User{
		Username: registerUser.Username,
		Name:     registerUser.Name,
		Email:    registerUser.Email,
		Password: registerUser.Password,
	}

	if err := database.DB.Create(&NewUser).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"succes": true,
		"data":    NewUser,
	})
}

func UserControllereditUser(ctx *fiber.Ctx) error {
	userid := ctx.Locals("user").(request.UserLogin).ID

	editUserRequest := new(request.UserUpdateRequest)
	var user entity.User

	if err := database.DB.Where("id =? ", userid).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find user",
			"error":   err,
		})
	}

	if err := ctx.BodyParser(editUserRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	if editUserRequest.Username != "" {
		user.Username = editUserRequest.Username

	}
	if editUserRequest.Name != "" {
		user.Name = editUserRequest.Name

	}
	if editUserRequest.Email != "" {
		user.Email = editUserRequest.Email

	}

	if err := database.DB.Save(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot update user",
			"error":   err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully update user",
		"data":    user,
	})

}

func UserControllerGetAll(ctx *fiber.Ctx) error {
	var user []entity.User

	err := database.DB.Preload("Acounts").Find(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot get user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success get all user",
		"data":    user,
	})
}
