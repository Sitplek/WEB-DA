package service

import (
    "WEB-DA/internal/models"
    "WEB-DA/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{storage: storage}
}

// Получить иерархию сотрудников
func (s *Service) GetEmployeeHierarchy(managerID, departmentID interface{}) ([]models.Employee, error) {
	// Передаем параметры в хранилище, где они будут обработаны
	return s.storage.GetEmployeeHierarchy(managerID, departmentID)
}


// Получить список всех сотрудников
func (s *Service) GetAllEmployees() ([]models.Employee, error) {
	return s.storage.GetAllEmployees()
}

// GetAllEmployeesWithFilters возвращает список сотрудников с фильтрами
func (s *Service) GetAllEmployeesWithFilters(filters map[string]interface{}) ([]models.Employee, error) {
	return s.storage.GetEmployeesWithFilters(filters)
}

// GetFilters возвращает список доступных фильтров
func (s *Service) GetFilters() ([]models.Filters, error) {
    return s.storage.GetFilters()
}

// GetSectionName возвращает название отдела по ID
func (s *Service) GetSectionName(id int) (string, error) {
    return s.storage.GetSectionName(id)
}