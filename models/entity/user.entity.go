package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Username  string `gorm:"unique" json:"username"`
	Name      string `json:"name"`
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	Acount    []Acount
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updatedat"`
	DeletedAt gorm.DeletedAt `json:"deletedat" gorm:"index"`
}
