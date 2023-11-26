package responses

import (

	"time"

	"gorm.io/gorm"
)


type Transaction struct {
	ID              uint64         `gorm:"primary_key:auto_increment" json:"id"`
	TransactionType string         `json:"transaction_type"`
	Amount          float64        `json:"amount"`
	Description     string         `json:"description"`
	AcountId        uint           `json:"acount_id"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updatedat"`
	DeletedAt       gorm.DeletedAt `json:"_" gorm:"index"`
}
