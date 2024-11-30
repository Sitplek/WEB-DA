	package handler

	import (
		"github.com/gin-gonic/gin"
		"WEB-DA/internal/service"
		"net/http"
	)

	type Handler struct {
		service *service.Service
	}

	func NewHandler(service *service.Service) *Handler {
		return &Handler{service: service}
	}

// Получение иерархии сотрудников
func (h *Handler) GetEmployeeHierarchy(c *gin.Context) {
    managerID := c.DefaultQuery("manager_id", "")  // получаем параметр, если он не передан - ставим пустую строку
    departmentID := c.DefaultQuery("department_id", "")

    // Если manager_id не пустой, парсим его в int
    var managerIDInt interface{}
    if managerID != "" {
        managerIDInt = managerID  // передаем строку, если она есть
    } else {
        managerIDInt = nil  // если пусто, передаем nil
    }

    // Если department_id не пустой, парсим его в int
    var departmentIDInt interface{}
    if departmentID != "" {
        departmentIDInt = departmentID  // передаем строку, если она есть
    } else {
        departmentIDInt = nil  // если пусто, передаем nil
    }

    // Вызов метода сервиса
    employees, err := h.service.GetEmployeeHierarchy(managerIDInt, departmentIDInt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, employees)
}


	// Получение списка всех сотрудников
	func (h *Handler) GetAllEmployees(c *gin.Context) {
		employees, err := h.service.GetAllEmployees()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, employees)
	}