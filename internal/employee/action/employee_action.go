package action

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/model"
	"github.com/sirupsen/logrus"
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
func (a *employeeAction) Create(ctx echo.Context, req model.Employee) error {
	logger := logrus.WithContext(ctx.Request().Context()).
		WithField("request", req)

	_, err := a.adminRepo.FindAll(ctx, map[string]interface{}{})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	return a.employeeRepo.Create(ctx, req)
}

// FindAll :nodoc:
func (a *employeeAction) FindAll(ctx echo.Context, req map[string]interface{}) (model.Employee, error) {
	logger := logrus.WithContext(ctx.Request().Context()).
		WithField("request", req)

	employees, err := a.employeeRepo.FindAll(ctx, req)
	if err != nil {
		logger.Error(err.Error())
		return model.Employee{}, err
	}
	return employees, nil
}
