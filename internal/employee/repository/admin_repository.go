package repository

import (
	"github.com/notblessy/model"
	"gorm.io/gorm"
)

type adminRepository struct {
	db *gorm.DB
}

// NewAdminRepository :nodoc:
func NewAdminRepository(db *gorm.DB) model.AdminRepository {
	return &adminRepository{
		db: db,
	}
}

// Create :nodoc:
func (a *adminRepository) Create(req model.Admin) error {
	return nil
}

// FindAll :nodoc:
func (a *adminRepository) FindAll(req map[string]interface{}) (model.Admin, error) {
	return model.Admin{}, nil
}
