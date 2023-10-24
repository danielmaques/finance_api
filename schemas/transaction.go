package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Add         bool
	Category    string
	Description string
	Amount      int64
	Date        time.Time
}
