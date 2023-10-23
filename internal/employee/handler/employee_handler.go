package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/notblessy/model"
	"github.com/notblessy/utils"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	employeeAction model.EmployeeAction
}

// NewHandler :nodoc:
func NewHandler(ea model.EmployeeAction) *Handler {
	return &Handler{
		employeeAction: ea,
	}
}

// FindAllHandler :nodoc:
func (t *Handler) FindAllHandler(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	employee, err := t.employeeAction.FindAll(c, map[string]interface{}{})
	if err != nil {
		logger.Error(err.Error())
		utils.Response(c, http.StatusInternalServerError, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}

	return utils.Response(c, http.StatusOK, &utils.HTTPResponse{
		Data: employee,
	})
}

// CreateHandler :nodoc:
func (t *Handler) CreateHandler(c echo.Context) error {
	logger := logrus.WithContext(c.Request().Context())

	err := t.employeeAction.Create(c, model.Employee{})
	if err != nil {
		logger.Error(err.Error())
		utils.Response(c, http.StatusInternalServerError, &utils.HTTPResponse{
			Message: err.Error(),
		})
	}
	return utils.Response(c, http.StatusOK, &utils.HTTPResponse{
		Data: model.Employee{},
	})
}
