package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"atehafinancetracker/database"
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/request"
	"atehafinancetracker/models/responses"
	trans "atehafinancetracker/models/responses/transaction"
)

// @Summary Create Transaction
// @Description Usage of this endpoint will Create Transaction on your account
// @Tags Transaction
// @Accept json
// @Produce json
// @Param TransactionBody body request.Transaction true "Transaction body"
// @Success 200 {object} trans.SuccessCreateTransaction "transaction  success"
// @Failure 400 {object} trans.FailureResponse  "transaction failed"
// @Router /transaction/create [post]
func TransactionControllerCreate(ctx *fiber.Ctx) error {
	transactionrequest := new(request.Transaction)
	var acount entity.Acount

	if err := ctx.BodyParser(transactionrequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(trans.FailureResponse{
			Success: false,
			Message: "Failed Create Transactiom",
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

	return ctx.Status(fiber.StatusCreated).JSON(trans.SuccessCreateTransaction{
		Data: newtransaction,
		Balance: int(acount.Balance),
		Message: "success",
		Success: true,
	})

}

// @Summary get all data by account id
// @Description Usage of this endpoint will get all data by account id
// @Tags Transaction
// @Accept json
// @Produce json
// @Param TransactionBody body request.Transaction true "Transaction body"
// @Success 200 {object} trans.SuccesGetTransaction "transaction  success"
// @Failure 400 {object} trans.FailureResponse  "transaction failed"
// @Router /transaction/{id} [get]
// @Param id path string true "id account to ger"
func TransactionControllerGetByAcountId(ctx *fiber.Ctx) error {

	var transaction []responses.Transaction
	acountid := ctx.Params("id")
	UserloginId := ctx.Locals("user").(request.UserLogin).ID

	err := database.DB.Table("transactions").Select("transactions.id, transactions.transaction_type, transactions.amount, transactions.description, transactions.acount_id, transactions.created_at, transactions.updated_at, transactions.deleted_at").Joins("JOIN acounts ON transactions.acount_id = acounts.id").Where("acounts.id = ? AND acounts.user_id =? ", acountid, UserloginId).Scan(&transaction).Error
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(trans.FailureResponse{
			Success: false,
			Message: err.Error(),
		})

	}
	if transaction == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Cannot find transaction",
		})
		
	}
	

	return ctx.Status(fiber.StatusOK).JSON(trans.SuccesGetTransaction{
		Success: true,
		Data: transaction,
		Message: "succes get data",
	})
}
