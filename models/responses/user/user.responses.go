package user

import (
	"atehafinancetracker/models/entity"
)

type SuccessRegister struct {
	Succes bool 		`json:"succes" example:"true" `
	Data entity.User	`json:"data"`
	Message string 		`json:"message" `
}

type FailureRegisterResponse struct {
	Succes bool 		`json:"succes" example:"false" `
	Message string 		`json:"message" example:"Email or Password is Already Exist" `
}

type GetAllUser struct {
	Succes bool 		`json:"succes" example:"true" `
	Data []entity.User	`json:"data"`
	Message string 		`json:"message" `
}
type FailureResponse struct {
	Succes bool 		`json:"succes" example:"false" `
	Message string 		`json:"message" example:"Can't get data" `
}

