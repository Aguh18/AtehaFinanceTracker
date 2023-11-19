package entity

import (
	"time"

	"gorm.io/gorm"
)

type Acount struct {
	ID          uint64  `gorm:"primary_key:auto_increment" json:"id"`
	Balance     float64 `json:"balance"`
	UserId      uint    `json:"user_id"`
	Transaction []Transaction
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updatedat"`
	DeletedAt   gorm.DeletedAt `json:"deletedat" gorm:"index"`
}
