package controller

import (
	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/models/responses"


	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func TransactionControllerCreate(ctx *fiber.Ctx) error {
	transactionrequest := new(request.Transaction)
	var acount entity.Acount

	if err := ctx.BodyParser(transactionrequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	validate := validator.New()

	if err := validate.Struct(transactionrequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation error",
			"error":   err.Error(),
		})
	}
	err := database.DB.First(&acount, "id =?", transactionrequest.AcountId).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find user",
			"error":   err,
		})
	}

	newtransaction := entity.Transaction{
		TransactionType: transactionrequest.TransactionType,
		Amount:          transactionrequest.Amount,
		Description:     transactionrequest.Description,
		AcountId:        transactionrequest.AcountId,
	}

	if transactionrequest.TransactionType == "DEBIT" {
		database.DB.Model(&acount).Update("balance", acount.Balance-transactionrequest.Amount)
	} else if transactionrequest.TransactionType == "CREDIT" {
		database.DB.Model(&acount).Update("balance", acount.Balance+transactionrequest.Amount)
	}

	if err := database.DB.Create(&newtransaction).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot create transaction",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Successfully created transaction",
		"data":    newtransaction,
		"balance": acount.Balance,
	})

}

func TransactionControllerGetByAcountId(ctx *fiber.Ctx) error {

	var transaction []responses.Transaction
	acountid := ctx.Params("id")
	UserloginId := ctx.Locals("user").(request.UserLogin).ID

	err := database.DB.Table("transactions").Select("transactions.id, transactions.transaction_type, transactions.amount, transactions.description, transactions.acount_id, transactions.created_at, transactions.updated_at, transactions.deleted_at").Joins("JOIN acounts ON transactions.acount_id = acounts.id").Where("acounts.id = ? AND acounts.user_id =? ", acountid, UserloginId).Scan(&transaction).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cannot find transaction",
			"error":   err,
		})

	}
	if transaction == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cannot find transaction",
		})
		
	}
	

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   transaction,
	})
}
