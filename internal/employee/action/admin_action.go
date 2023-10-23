package action

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/model"
	"github.com/sirupsen/logrus"
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
func (a *adminAction) Create(c echo.Context, req model.Admin) error {
	err := a.repository.Create(c, req)
	if err != nil {
		logrus.WithContext(c.Request().Context()).
			WithField("request", req).
			Error(err.Error())
		return err
	}

	return nil
}

// FindAll :nodoc:
func (a *adminAction) FindAll(c echo.Context, req map[string]interface{}) (model.Admin, error) {
	admins, err := a.repository.FindAll(c, req)
	if err != nil {
		logrus.WithContext(c.Request().Context()).
			WithField("request", req).
			Error(err.Error())
		return model.Admin{}, err
	}
	return admins, nil
}
