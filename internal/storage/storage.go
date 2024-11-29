package storage

import (
    "gorm.io/gorm"
    "WEB-DA/internal/models"
)

type OrganizationStorage struct {
    db *gorm.DB
}

func NewOrganizationStorage(db *gorm.DB) *OrganizationStorage {
    return &OrganizationStorage{db: db}
}

// Получение дерева организации
func (s *OrganizationStorage) GetAllDepartments() ([]models.Department, error) {
    var departments []models.Department
    err := s.db.Preload("Employees.Supervisor").Find(&departments).Error
    return departments, err
}

// Фильтрация сотрудников
func (s *OrganizationStorage) GetFilteredEmployees(filters map[string]interface{}) ([]models.Employee, error) {
    var employees []models.Employee
    query := s.db.Model(&models.Employee{})
    for key, value := range filters {
        query = query.Where(key+" = ?", value)
    }
    err := query.Find(&employees).Error
    return employees, err
}
