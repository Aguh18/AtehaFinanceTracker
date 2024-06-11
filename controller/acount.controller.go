package controller

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/models/responses"
	"atehafinancetracker/models/responses/account"
)

// @Summary Create Account
// @Description Usage of this endpoint is for Create new Account
// @Tags Account
// @Accept json
// @Produce json
// @Param CreateBody body request.AcountRequestCreate true "Create Account Data"
// @Success 201 {object} account.SuccesCreateAccountResponses "Succes Create"
// @Failure 400 {object} account.FailureResponse "Failed Create"
// @Param x-token header string true "Jwt Token"
// @Router /acount/create [post]
func AcountControllerCreate(ctx *fiber.Ctx) error {
	acountcreateRequest := new(request.AcountRequestCreate)
	user := ctx.Locals("user").(request.UserLogin)

	if err := ctx.BodyParser(acountcreateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(account.FailureResponse{
			Success: false,
			Message: "Cannot parse JSON",
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

	return ctx.Status(fiber.StatusCreated).JSON(account.SuccesCreateAccountResponses{
		Success: true,
		Data:   newAcount,
		Message: "Succes Create Account",
	})

}

// @Summary Get all Account by user Id
// @Description Usage of this endpoint is for Create new Account
// @Tags Account
// @Accept json
// @Produce json
// @Success 200 {object} account.SuccesGetDataResponses "Succes Get all data By id"
// @Failure 400 {object} account.FailureResponse "Failed get data"
// @Param x-token header string true "Jwt Token"
// @Router /acount [get]
func AcountControllerGetByUserId(ctx *fiber.Ctx) error {

	Userid := ctx.Locals("user").(request.UserLogin).ID
	var acount []responses.Acount

	err := database.DB.Where("user_id =?", Userid).Find(&acount).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(account.FailureResponse{
			Success: false,
			Message: "Can't get data",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(account.SuccesGetDataResponses{
		Success: true,
		Data:    acount,
		Message: "Succes Get Data",
	})

}


// @Summary Update Data Account
// @Description Usage of this endpoint is for Update Data Account by id
// @Tags Account
// @Accept json
// @Produce json
// @Param CreateBody body request.AcountRequestUpdate true "Update Data"
// @Success 200 {object} account.SuccessUpdateAccountResponses "Succes Get all data By id"
// @Failure 400 {object} account.FailureResponse "Failed get data"
// @Param x-token header string true "Jwt Token"
// @Router /acount/{id} [put]
// @Param id path string true "id user to update"]
func AcountControllerUpdate(ctx *fiber.Ctx) error {

	userId := ctx.Locals("user").(request.UserLogin).ID
	log.Println(userId)

	var acount entity.Acount

	acountUpdateRequest := new(request.AcountRequestUpdate)

	if err := ctx.BodyParser(acountUpdateRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(account.FailureResponse{
			Success: false,
			Message: "Cannot parse JSON",
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

	return ctx.Status(fiber.StatusOK).JSON(account.SuccessUpdateAccountResponses{
		Success: true,
		Message: "Succes Update Data",
		Data:    acount,
	})
}



// @Summary Delete account By id
// @Description Usage of this endpoint is for delete account by id
// @Tags Account
// @Accept json
// @Produce json
// @Success 200 {object} account.SuccessUpdateAccountResponses "Succes Get all data By id"
// @Failure 400 {object} account.FailureResponse "Failed get data"
// @Param x-token header string true "Jwt Token"
// @Router /acount/{id} [delete]
// @Param id path string true "id user to update"
func AcountControllerDeleteById(ctx *fiber.Ctx) error {

	param := ctx.Params("id")
	userId := ctx.Locals("user").(request.UserLogin).ID
	Acount := new(entity.Acount)
	if param == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(account.FailureResponse{
			Success: false,
			Message: "Acount id is required",
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

	return ctx.Status(fiber.StatusOK).JSON(account.SuccessResponse{
		Success: true,
		Message: "Succes Delete Data",
	})
	
}
