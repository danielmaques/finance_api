package handler

import (
	"fmt"
	"time"
)

func errParamsRequired(name, typ string) error {
	return fmt.Errorf("parameter %s is required of type %s", name, typ)
}

type CreateTransactionRequest struct {
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

	if r.Category == "" {
		return errParamsRequired("category", "string")
	}

	if r.Description == "" {
		return errParamsRequired("description", "string")
	}

	if r.Amount <= 0 {
		return errParamsRequired("amount", "int64")
	}

	if r.Date.IsZero() {
		return errParamsRequired("date", "time.Time")
	}

	return nil
}
