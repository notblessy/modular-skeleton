package repository

import (
	"github.com/notblessy/model"
	"gorm.io/gorm"
)

type employeeRepository struct {
	db *gorm.DB
}

// NewEmployeeRepository :nodoc:
func NewEmployeeRepository(db *gorm.DB) model.EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

// Create :nodoc:
func (a *employeeRepository) Create(req model.Employee) error {
	return nil
}

// FindAll :nodoc:
func (a *employeeRepository) FindAll(req map[string]interface{}) (model.Employee, error) {
	return model.Employee{}, nil
}
