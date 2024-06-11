package request

type User struct {
	Username string `json:"username" validate:"required" example:"guhh"`
	Name     string `json:"name" validate:"required" example:"guhh"`
	Email    string `json:"email" validate:"required,email" example:"teguh22@gmail.com"`
	Password string `json:"password" validate:"required" example:"password123"`
}

type UserUpdateRequest struct {
	Username string `json:"username" `
	Name     string `json:"name" `
	Email    string `json:"email" `
}