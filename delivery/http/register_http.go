package http

import "github.com/notblessy/internal/employee"

// RegisterEmployeeTransport :nodoc:
func (h *HTTPService) RegisterEmployeeTransport(et *employee.Transport) {
	h.employee = *et
}
