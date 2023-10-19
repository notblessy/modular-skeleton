package employee

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/internal/employee/action"
	"github.com/notblessy/internal/employee/handler"
	"github.com/notblessy/internal/employee/repository"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Transport struct {
	db     *gorm.DB
	tracer trace.Tracer
}

// NewTransport :nodoc:
func NewTransport(db *gorm.DB, trace trace.Tracer) *Transport {
	return &Transport{
		db:     db,
		tracer: trace,
	}
}

// NewTransport :nodoc:
func (t *Transport) Transport(routes *echo.Group) {
	h := t.initiateEmployee()

	routes.POST("/one", h.CreateHandler)
	routes.GET("/list", h.FindAllHandler)
}

func (t *Transport) initiateEmployee() *handler.Handler {
	employeeRepo := repository.NewEmployeeRepository(t.db)
	adminRepo := repository.NewAdminRepository(t.db)

	employeeAction := action.NewEmployeeAction(employeeRepo, adminRepo)

	return handler.NewHandler(employeeAction)
}
