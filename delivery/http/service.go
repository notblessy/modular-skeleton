package http

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/internal/employee"
)

type HTTPService struct {
	employee employee.Transport
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// Router :nodoc:
func (h *HTTPService) Router(router *echo.Echo) {
	v1 := router.Group("/v1")

	h.employee.Transport(v1.Group("/employees"))
}
