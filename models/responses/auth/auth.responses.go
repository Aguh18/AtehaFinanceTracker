package auth


type SuccesResponse struct {
	Message string `json:"message" example:"Login Succes" `
	Success bool   `json:"success" example:"true"  `
	Token   string `json:"token" example:"jhdfkjsbgakjhgsgkhksjgb" `
}
type FailureResponse struct {
	Message string `json:"message" example:"Email or Password is wrong" `
	Success bool   `json:"success" example:"false"  `
}