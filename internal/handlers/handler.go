package handlers

import (
    "github.com/gin-gonic/gin"
    "WEB-DA/internal/service"
    "net/http"
)

type OrganizationHandler struct {
    service *service.OrganizationService
}

func NewOrganizationHandler(service *service.OrganizationService) *OrganizationHandler {
    return &OrganizationHandler{service: service}
}

// Endpoint для дерева организации
func (h *OrganizationHandler) GetOrganizationTree(c *gin.Context) {
    tree, err := h.service.GetOrganizationTree()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tree)
}

// Endpoint для фильтрации сотрудников
func (h *OrganizationHandler) GetFilteredEmployees(c *gin.Context) {
    filters := map[string]interface{}{}
    for _, key := range []string{"name", "role", "job_position", "email", "phone"} {
        if value := c.Query(key); value != "" {
            filters[key] = value
        }
    }
    employees, err := h.service.GetFilteredEmployees(filters)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, employees)
}
