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
// @Description Возвращает список всех сотрудников с возможностью фильтрации
// @Tags Employees
// @Produce json
// @Param first_name query string false "Имя"
// @Param last_name query string false "Фамилия"
// @Param position query string false "Должность"
// @Param role query string false "Роль"
// @Param phone query string false "Телефон"
// @Param email query string false "Почта"
// @Param manager_id query int false "ID руководителя"
// @Param department_id query int false "ID департамента"
// @Param division_id query int false "ID подразделения"
// @Param unit_id query int false "ID отдела"
// @Success 200 {array} models.Employee
// @Failure 500 {object} ErrorResponse
// @Router /employees [get]
func (h *Handler) GetAllEmployees(c *gin.Context) {
	// Собираем фильтры из параметров запроса
	filters := map[string]interface{}{
		"first_name":   c.Query("first_name"),
		"last_name":    c.Query("last_name"),
		"photo_path":	c.Query("photo_path"),
		"position":     c.Query("position"),
		"role":         c.Query("role"),
		"phone":        c.Query("phone"),
		"email":        c.Query("email"),
		"manager_id":   c.Query("manager_id"),
		"department_id": c.Query("department_id"),
		"division_id":  c.Query("division_id"),
		"unit_id":      c.Query("unit_id"),
	}

	// Вызов слоя сервиса с фильтрами
	employees, err := h.service.GetAllEmployeesWithFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// GetFilters godoc
// @Summary Получение списка доступных фильтров
// @Description Возвращает список фильтров для сотрудников
// @Tags Employees
// @Produce json
// @Success 200 {array} models.Filters
// @Failure 500 {object} ErrorResponse
// @Router /filters [get]
func (h *Handler) GetFilters(c *gin.Context) {
    filters, err := h.service.GetFilters()
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, filters)
}

