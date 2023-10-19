package model

// EmployeeRepository :nodoc:
type EmployeeRepository interface {
	Create(employee Employee) error
	FindAll(query map[string]interface{}) (Employee, error)
}

// EmployeeAction :nodoc:
type EmployeeAction interface {
	Create(employee Employee) error
	FindAll(query map[string]interface{}) (Employee, error)
}

// Employee :nodoc:
type Employee struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
