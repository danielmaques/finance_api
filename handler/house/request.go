package house

import (
	"fmt"

	"github.com/danielmaques/finance_api/schemas"
)

func errParamsHouseRequired(name, typ string) error {
	return fmt.Errorf("parameter %s is required of type %s", name, typ)
}

type CreateHouseRequest struct {
	InviteCode string         `json:"invite_code"`
	Users      []schemas.User `json:"users"`
}

func (r *CreateHouseRequest) Validate() error {
	if len(r.Users) == 0 {
		return errParamsHouseRequired("users", "[]User")
	}

	return nil
}
