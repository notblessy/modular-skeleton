package model

import "github.com/labstack/echo/v4"

// EmployeeRepository :nodoc:
type EmployeeRepository interface {
	Create(ctx echo.Context, employee Employee) error
	FindAll(ctx echo.Context, query map[string]interface{}) (Employee, error)
}

// EmployeeAction :nodoc:
type EmployeeAction interface {
	Create(ctx echo.Context, employee Employee) error
	FindAll(ctx echo.Context, query map[string]interface{}) (Employee, error)
}

// Employee :nodoc:
type Employee struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
