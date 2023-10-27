package schemas

import (
	"time"

	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	Terms        bool
	Users        []User        `gorm:"foreignKey:HouseID"`
	Transactions []Transaction `gorm:"foreignKey:HouseID"`
}

type HouseResponse struct {
	ID          uint64        `json:"id"`
	CreatedAt   time.Time     `json:"created_at"`
	Balance     float64       `json:"balance"`
	Users       []User        `json:"user"`
	Transaction []Transaction `json:"transaction"`
}
