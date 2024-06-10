package request



type LoginRequest struct {
	Email    string `json:"email" validate:"required,email" example:"conthjoh@emle.com"`
	Password string `json:"password" validate:"required" example:"conthgjoh_passrd"`
}

type UserLogin struct {
	ID       float64
	Username string
	Email    string
}
