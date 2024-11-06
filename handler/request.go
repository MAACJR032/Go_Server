package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func errParamRequired(name, typ string) error {
    return fmt.Errorf("parameter: %s (type: %s) is required", name, typ)
}

func (r *CreateOpeningRequest) Validate() error {
    if r.Role == "" && r.Company == "" && r.Location == "" && 
       r.Link == "" && r.Remote == nil && r.Salary <= 0 {
        return fmt.Errorf("request body is empty or malformed")
    }
	if r.Role == "" {
		return errParamRequired("role", "string")
	}
    if r.Company == "" {
		return errParamRequired("company", "string")
	}
    if r.Location == "" {
        return errParamRequired("location", "string")
    }
    if r.Link == "" {
		return errParamRequired("link", "string")
	}
    if r.Remote == nil {
		return errParamRequired("Remote", "bool")
	}
    if r.Salary <= 0 {
		return errParamRequired("salary", "int64")
	}

    return nil
}