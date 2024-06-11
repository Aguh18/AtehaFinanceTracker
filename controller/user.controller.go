package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"

	usr "atehafinancetracker/models/responses/user"
	"atehafinancetracker/utils"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

// @Summary Register User
// @Description Usage of this endpoint will Create User For Login
// @Tags User
// @Accept json
// @Produce json
// @Param loginRequest body request.User true "User login information"
// @Success 200 {object} usr.SuccessRegister "Login success"
// @Failure 400 {object} usr.FailureRegisterResponse "Invalid request"
// @Router /register [post]
func UserControllerRegister(ctx *fiber.Ctx) error {

	registerUser := new(request.User)

	if err := ctx.BodyParser(registerUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(usr.FailureRegisterResponse{
			Succes: false,
			Message: "Cannot parse JSON",

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

	return ctx.Status(fiber.StatusOK).JSON(usr.SuccessRegister{
		Succes: true,
		Data: NewUser,
	})
}

// @Summary Update User
// @Description Usage of this endpoint is for gett all user without jwt param
// @Tags User
// @Accept json
// @Produce json
// @Param UpdateRequest body request.UserUpdateRequest true "User login information"
// @Param x-token header string true "Jwt Token"
// @Success 200 {object} usr.GetAllUser "Succes Update Data"
// @Failure 400 {object} usr.FailureResponse "Update data failed"
// @Router /user/{id} [get]
// @Param id path string true "id user to update"
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

// @Summary Get All data
// @Description Usage of this endpoint is for gett all user without jwt param
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} usr.GetAllUser "Succes Get All Data"
// @Failure 400 {object} usr.FailureResponse "Invalid request"
// @Router /user [get]
func UserControllerGetAll(ctx *fiber.Ctx) error {
	var user []entity.User

	err := database.DB.Preload("Acounts").Find(&user).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(usr.FailureResponse{
			Succes: false,
			Message: "Can't get data",
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(usr.GetAllUser{
		Succes: true,
		Data:  user,
		Message: "Successfully get all user",
	})
}
