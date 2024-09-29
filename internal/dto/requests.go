package dto

import "fmt"

type SigninRequest struct {
	Email    string
	Password string
}

type SignupRequest struct {
	Email     string `json: "email"`
	LastName  string `json: "lastName"`
	FirstName string `json: "firstName"`
	Password  string `json: "password"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *SignupRequest) Validate() error {
	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}
	if r.LastName == "" {
		return errParamIsRequired("LastName", "string")
	}
	if r.FirstName == "" {
		return errParamIsRequired("FirstName", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("Password", "string")
	}

	return nil
}

func (r *SigninRequest) Validate() error {
	if r.Email == "" {
		return errParamIsRequired("Email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("Password", "string")
	}

	return nil
}
