package house

import "fmt"

func errParamsHouseRequired(name, typ string) error {
	return fmt.Errorf("parameter %s is required of type %s", name, typ)
}

type CreateHouseRequest struct {
	Terms bool `json:"terms"`
}

func (r *CreateHouseRequest) Validate() error {
	if r.Terms == false {
		return errParamsHouseRequired("terms", "bool")
	}

	return nil
}
