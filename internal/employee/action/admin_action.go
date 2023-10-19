package action

import (
	"github.com/notblessy/model"
)

type adminAction struct {
	repository model.AdminRepository
}

// NewAdminAction :nodoc:
func NewAdminAction(repo model.AdminRepository) model.AdminAction {
	return &adminAction{
		repository: repo,
	}
}

// Create :nodoc:
func (a *adminAction) Create(req model.Admin) error {
	return a.repository.Create(req)
}

// FindAll :nodoc:
func (a *adminAction) FindAll(req map[string]interface{}) (model.Admin, error) {
	return a.repository.FindAll(req)
}
