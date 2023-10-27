package schemas

import (
	"time"

	"gorm.io/gorm"
)

type House struct {
	gorm.Model
	InviteCode string
	Users      []User `gorm:"foreignKey:HouseID"`
}

type HouseResponse struct {
	ID          uint64        `json:"id"`
	InviteCode  string        `json:"invite_code"`
	CreatedAt   time.Time     `json:"created_at"`
	Balance     float64       `json:"balance"`
	Users       []User        `json:"user"`
	Transaction []Transaction `json:"transaction"`
}
