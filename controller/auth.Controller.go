package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/utils"
	
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(ctx *fiber.Ctx) error {
	user := new(entity.User)
	LoginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(LoginRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	validate := validator.New()

	if err := validate.Struct(LoginRequest); err != nil {

		ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	err := database.DB.Debug().First(&user, "email = ?", LoginRequest.Email).Error
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Email or Password is wrong",
		})

	}

	isValid := utils.CompareHashedPassword(LoginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Email or Password is wrong",
		})
	}


	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	tokenString , err := utils.GenerateToken(&claims)
	if err != nil{

		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}



	return ctx.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Login Success",
		"token": tokenString,

	})

}
