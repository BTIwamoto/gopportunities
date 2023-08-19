package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (c *CreateOpeningRequest) Validate() error {
	if c.Role == "" && c.Company == "" && c.Location == "" && c.Remote == nil && c.Link == "" && c.Salary < 1 {
		return fmt.Errorf("request body is empty")
	}
	if c.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if c.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if c.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if c.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if c.Salary < 1 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

// UpdateOpening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (u *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is truthy
	if u.Role != "" || u.Company != "" || u.Location != "" || u.Remote != nil || u.Link != "" || u.Salary > 0 {
		return nil
	}

	// if none oof the fields werer provided, return falsy
	return fmt.Errorf("at least one valid field must be provided")
}
