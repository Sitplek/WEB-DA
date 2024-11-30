package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"WEB-DA/internal/service"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Handler — структура обработчика запросов.
type Handler struct {
	service *service.Service
}

// NewHandler создает новый экземпляр обработчика.
func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// GetEmployeeHierarchy godoc
// @Summary Получение иерархии сотрудников
// @Description Возвращает иерархию сотрудников по manager_id или department_id
// @Tags Employees
// @Produce json
// @Param manager_id query int false "ID руководителя"
// @Param department_id query int false "ID отдела"
// @Success 200 {array} models.Employee
// @Failure 500 {object} ErrorResponse
// @Router /employees/hierarchy [get]
func (h *Handler) GetEmployeeHierarchy(c *gin.Context) {
	// Получение параметров запроса
	managerID := c.DefaultQuery("manager_id", "")
	departmentID := c.DefaultQuery("department_id", "")

	// Преобразование параметров в интерфейсы
	var managerIDInt, departmentIDInt interface{}
	if managerID != "" {
		managerIDInt = managerID
	}
	if departmentID != "" {
		departmentIDInt = departmentID
	}

	// Вызов метода сервиса
	employees, err := h.service.GetEmployeeHierarchy(managerIDInt, departmentIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ответ
	c.JSON(http.StatusOK, employees)
}

// GetAllEmployees godoc
// @Summary Получение всех сотрудников
// @Description Возвращает список всех сотрудников
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Employee
// @Failure 500 {object} ErrorResponse
// @Router /employees [get]
func (h *Handler) GetAllEmployees(c *gin.Context) {
	// Вызов метода сервиса
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Ответ
	c.JSON(http.StatusOK, employees)
}
