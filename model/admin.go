package model

// AdminRepository :nodoc:
type AdminRepository interface {
	Create(admin Admin) error
	FindAll(query map[string]interface{}) (Admin, error)
}

// AdminAction :nodoc:
type AdminAction interface {
	Create(admin Admin) error
	FindAll(query map[string]interface{}) (Admin, error)
}

// Admin :nodoc:
type Admin struct {
	Id        string `json:"id"`
	CompanyId string `json:"companyId"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}
