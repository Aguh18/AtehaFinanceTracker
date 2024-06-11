package account

import (
	"atehafinancetracker/models/entity"
	"atehafinancetracker/models/responses"
)

type FailureResponse struct {
	Message string `json:"message" example:"Failed create Account" `
	Success bool   `json:"success" example:"false"  `
}

type SuccessResponse struct {
	Message string `json:"message" example:"succes" `
	Success bool   `json:"success" example:"true"  `
}
type SuccesCreateAccountResponses struct {
	Message string `json:"message" example:"Succes Create Account" `
	Success bool   `json:"success" example:"true"  `
	Data  entity.Acount `json:"data" `
}

type SuccesGetDataResponses struct {
	Message string `json:"message" example:"Succes Create Account" `
	Success bool   `json:"success" example:"true"  `
	Data  []responses.Acount `json:"data" `
}

type SuccessUpdateAccountResponses struct {
	Success bool 		`json:"succes" example:"true" `
	Message string 		`json:"message" example:"succes update Data" `
	Data entity.Acount `json:"data" ` 
}

