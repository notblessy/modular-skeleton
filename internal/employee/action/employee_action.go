package action

import (
	"github.com/notblessy/model"
)

type employeeAction struct {
	employeeRepo model.EmployeeRepository
	adminRepo    model.AdminRepository
}

// NewEmployeeAction :nodoc:
func NewEmployeeAction(er model.EmployeeRepository, ar model.AdminRepository) model.EmployeeAction {
	return &employeeAction{
		employeeRepo: er,
		adminRepo:    ar,
	}
}

// Create :nodoc:
func (a *employeeAction) Create(req model.Employee) error {
	_, err := a.adminRepo.FindAll(map[string]interface{}{})
	if err != nil {
		return err
	}

	return a.employeeRepo.Create(req)
}

// FindAll :nodoc:
func (a *employeeAction) FindAll(req map[string]interface{}) (model.Employee, error) {
	return a.employeeRepo.FindAll(req)
}
