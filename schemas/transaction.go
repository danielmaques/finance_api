package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	HouseID     uint64 `gorm:"not null"`
	Add         bool
	Category    string
	Description string
	Amount      int64
	Date        time.Time
}

type TransactionResponse struct {
	ID          uint64    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Add         bool      `json:"add"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}
