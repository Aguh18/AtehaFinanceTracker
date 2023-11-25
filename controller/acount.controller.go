package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/models/responses"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func AcountControllerCreate(ctx *fiber.Ctx) error {
	acountcreateRequest := new(request.AcountRequestCreate)
	user := ctx.Locals("user").(request.UserLogin)

	if err := ctx.BodyParser(acountcreateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}
	validate := validator.New()

	if err := validate.Struct(acountcreateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err,
		})
	}

	newAcount := entity.Acount{
		AcountName: acountcreateRequest.AcountName,
		Balance:    acountcreateRequest.Balance,
		UserId:     uint(user.ID),
	}

	if err := database.DB.Create(&newAcount).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create acount",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{

		"message": "Successfully created acount",
		"data":    newAcount,
	})

}

func AcountControllerGetByUserId(ctx *fiber.Ctx) error {

	Userid := ctx.Locals("user").(request.UserLogin).ID
	var acount []responses.Acount

	err := database.DB.Where("user_id =?", Userid).Find(&acount).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot find acount",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   acount,
	})

}

func AcountControllerUpdate(ctx *fiber.Ctx) error {

	userId := ctx.Locals("user").(request.UserLogin).ID
	log.Println(userId)

	var acount entity.Acount

	acountUpdateRequest := new(request.AcountRequestUpdate)

	if err := ctx.BodyParser(acountUpdateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}

	param := ctx.Params("id")
	if param == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Acount id is required",
		})
	}

	if err := database.DB.Where("id = ? AND user_id = ? ", param, userId).First(&acount).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot find acount",
			"error":   err,
		})
	}

	if acountUpdateRequest.AcountName != "" {
		acount.AcountName = acountUpdateRequest.AcountName
	}

	if acountUpdateRequest.Balance != 0 {
		acount.Balance = acountUpdateRequest.Balance

	}

	if err := database.DB.Save(&acount).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot update acount",
			"error":   err,
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully update acount",
		"data":    acount,
	})
}



func AcountControllerDeleteById(ctx *fiber.Ctx) error {

	param := ctx.Params("id")
	userId := ctx.Locals("user").(request.UserLogin).ID
	Acount := new(entity.Acount)
	if param == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Acount id is required",
		})
		
	}
	err := database.DB.Where("id = ? AND user_id = ?",param, userId).Delete(&Acount).Error

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot delete acount",
			"error":   err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully delete acount",
		
	})
	
}
