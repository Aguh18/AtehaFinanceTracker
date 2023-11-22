package request


type AcountRequestCreate struct {
	AcountName string `json:"acount_name" validate:"required"`
	Balance    float64 `json:"balance" validate:"required"`
	
	
}