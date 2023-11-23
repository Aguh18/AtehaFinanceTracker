package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)




func AcountControllerCreate(ctx *fiber.Ctx) error {	
	acountcreateRequest := new(request.AcountRequestCreate)
	user := ctx.Locals("user").(request.UserLogin)


	if err := ctx.BodyParser(acountcreateRequest); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	validate := validator.New()

	if err := validate.Struct(acountcreateRequest); err != nil{
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error": err,
		})
	}
	

	newAcount := entity.Acount{
		AcountName: acountcreateRequest.AcountName,
		Balance: acountcreateRequest.Balance,
		UserId: uint(user.ID),
	}

	if err := database.DB.Create(&newAcount).Error; err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create acount",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		
		"message": "Successfully created acount",
		"data": newAcount,
	})

}


func AcountControllerGetByUserId(ctx *fiber.Ctx) error  {
	
	Userid := ctx.Params("id")
	var acount []entity.Acount

	if Userid == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Userid is required",
		})
	}

	err := database.DB.Where("user_id =?",Userid).Find(&acount).Error
	if  err != nil  {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": "Cannot find acount",
			"error": err,
		})
	}


	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": acount,
	})


}