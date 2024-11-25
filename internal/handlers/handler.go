package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"WEB-DA/internal/models"
	"WEB-DA/internal/service"
)

type Handlers struct {
	service service.Service
}

func NewHandlers(srv service.Service) *Handlers {
	return &Handlers{service: srv}
}

// GetOrganizationTree возвращает дерево организаций
func (h *Handlers) GetOrganizationTree(c *gin.Context) {
	tree, err := h.service.GetOrganizationTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tree)
}

// EmployeeSearch ищет сотрудников по фильтрам
func (h *Handlers) EmployeeSearch(c *gin.Context) {
	position := c.Query("position")
	skill := c.Query("skill")
	departmentID := c.Query("department_id")

	var depID *int
	if departmentID != "" {
		id, err := strconv.Atoi(departmentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department_id"})
			return
		}
		depID = &id
	}

	employees, err := h.service.SearchEmployees(position, skill, depID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

// GetEmployeeDetails возвращает подробности о сотруднике
func (h *Handlers) GetEmployeeDetails(c *gin.Context) {
	idParam := c.Param("id")
	employeeID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	employee, err := h.service.GetEmployeeDetails(employeeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}

// UpdateEmployee обновляет информацию о сотруднике
func (h *Handlers) UpdateEmployee(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.UpdateEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
}

// AddDepartment добавляет новый департамент
func (h *Handlers) AddDepartment(c *gin.Context) {
	var department models.Department
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.AddDepartment(department); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Department added successfully"})
}
