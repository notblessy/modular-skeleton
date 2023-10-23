package model

import "github.com/labstack/echo/v4"

// AdminRepository :nodoc:
type AdminRepository interface {
	Create(ctx echo.Context, admin Admin) error
	FindAll(ctx echo.Context, query map[string]interface{}) (Admin, error)
}

// AdminAction :nodoc:
type AdminAction interface {
	Create(ctx echo.Context, admin Admin) error
	FindAll(ctx echo.Context, query map[string]interface{}) (Admin, error)
}

// Admin :nodoc:
type Admin struct {
	Id        string `json:"id"`
	CompanyId string `json:"companyId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
