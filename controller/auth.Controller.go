package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/utils"

	"atehafinancetracker/models/responses/auth"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// @Summary Login User
// @Description Usage of this endpoint will allow user to login and get JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "User login information"
// @Success 200 {object} auth.SuccesResponse "Login success"
// @Failure 400 {object} auth.FailureResponse "Invalid request"
// @Router /login [post]
// @Example loginSuccessExample
// @Example loginBadRequestExample
func Login(ctx *fiber.Ctx) error {
	user := new(entity.User)
	LoginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(LoginRequest); err != nil {
		
		
		return ctx.Status(400).JSON(auth.FailureResponse{
			Message: "Invalid request",
			Success: false,
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
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["Id"] = user.ID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString , err := utils.GenerateToken(&claims)
	if err != nil{

		return ctx.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}


	
	return ctx.Status(200).JSON(auth.SuccesResponse{
		Message: "Login Succes",
		Success: true,
		Token: tokenString,
	})
}
