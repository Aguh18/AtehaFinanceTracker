package transaction

import (
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/responses"
)

type FailureResponse struct {
	Message string `json:"message" example:"Failed create " `
	Success bool   `json:"success" example:"false"  `
}

type SuccessResponse struct {
	Message string `json:"message" example:"succes" `
	Success bool   `json:"success" example:"true"  `
}


type SuccessCreateTransaction struct {
	Message string `json:"message" example:"succes" `
	Success bool   `json:"success" example:"true"  `
	Data entity.Transaction `json:"data" `
	Balance int `json:"balance" example:"200000" `


}

type SuccesGetTransaction struct {
	Message string `json:"message" example:"succes" `
	Success bool   `json:"success" example:"true"  `
	Data []responses.Transaction `json:"data" `



}