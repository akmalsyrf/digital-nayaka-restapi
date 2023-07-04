package handler

import (
	"digital-nayaka-test/employee"

	"github.com/gin-gonic/gin"
)

type employeeHandler struct {
	employeeService employee.Service
}

func NewEmployeeHandler(employeeService employee.Service) *employeeHandler {
	return &employeeHandler{employeeService}
}

func (h *employeeHandler) GetEmployees(c *gin.Context) {
	// userID, _ := strconv.Atoi(c.Query("user_id"))

	employees, err := h.employeeService.GetEmployees()
	if err != nil {
		c.JSON(400, employees)
		return
	}

	c.JSON(200, employees)
}
