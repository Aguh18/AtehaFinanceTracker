package request

type Transaction struct {
	TransactionType string  `json:"transaction_type" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
	Description     string  `json:"description" ` 
	AcountId        uint    `json:"acount_id" validate:"required"`
}
