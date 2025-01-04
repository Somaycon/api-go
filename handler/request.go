package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)

}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Name == "" && r.Description == "" {
		return fmt.Errorf("Boby is empty")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

type UpdateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *UpdateTaskRequest) Validate() error {
	if r.Name != "" || r.Description != "" {
		return nil
	}
	return fmt.Errorf("at least one valid field request must be provided")
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Email == "" && r.Password == "" {
		return fmt.Errorf("Body is empty")
	}
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	return nil
}
