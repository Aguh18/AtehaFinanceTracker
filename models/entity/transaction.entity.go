package entity

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type transactionType string

const (
	DEBIT  transactionType = "DEBIT"
	CREDIT transactionType = "CREDIT"
)

func (tt *transactionType) Scan(value interface{}) error {
	*tt = transactionType(value.([]byte))
	return nil
}
func (tt transactionType) Value() (driver.Value, error) {
	return string(tt), nil
}

type Transaction struct {
	ID              uint64          `gorm:"primary_key:auto_increment" json:"id"`
	TransactionType transactionType `gorm:"type:transaction_type"`
	Amount          float64         `json:"amount"`
	Description     string          `json:"description"`
	AcountId        uint            `json:"acount_id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updatedat"`
	DeletedAt       gorm.DeletedAt  `json:"deletedat" gorm:"index"`
}
