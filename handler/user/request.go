package user

import "fmt"

func errParamsUserRequired(name, typ string) error {
	return fmt.Errorf("parameter %s is required of type %s", name, typ)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" || r.Email == "" || r.Password == "" {
		return errParamsUserRequired("name", "string")
	}

	if r.Email == "" {
		return errParamsUserRequired("email", "string")
	}

	if r.Password == "" {
		return errParamsUserRequired("password", "string")
	}

	return nil
}
