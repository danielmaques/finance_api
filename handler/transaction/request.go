package transaction

import (
	"fmt"
	"time"
)

func errParamsTransactionRequired(name, typ string) error {
	return fmt.Errorf("parameter %s is required of type %s", name, typ)
}

type CreateTransactionRequest struct {
	UserID      uint64      `json:"user_id"`
	Add         bool      `json:"add"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	Date        time.Time `json:"date"`
}

func (r *CreateTransactionRequest) Validate() error {

	if r.Category == "" && r.Description == "" && r.Amount == 0 && r.Date.IsZero() {
		return fmt.Errorf("request body is empty")
	}

	if r.UserID == 0 {
		return errParamsTransactionRequired("user_id", "uint64")
	}

	if r.Category == "" {
		return errParamsTransactionRequired("category", "string")
	}

	if r.Description == "" {
		return errParamsTransactionRequired("description", "string")
	}

	if r.Amount <= 0 {
		return errParamsTransactionRequired("amount", "int64")
	}

	if r.Date.IsZero() {
		return errParamsTransactionRequired("date", "time.Time")
	}

	return nil
}
