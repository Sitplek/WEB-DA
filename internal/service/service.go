package service

import (
    "WEB-DA/internal/models"
    "WEB-DA/internal/storage"
)

type Service interface {
    GetOrganizationTree() ([]models.Department, error)
    SearchEmployees(position, skill string, departmentID *int) ([]models.Employee, error)
    GetEmployeeDetails(employeeID int) (*models.Employee, error)
    UpdateEmployee(employee models.Employee) error
    AddDepartment(department models.Department) error
}

type serviceImpl struct {
    storage storage.Storage
}

func NewService(storage storage.Storage) Service {
    return &serviceImpl{storage: storage}
}

func (s *serviceImpl) GetOrganizationTree() ([]models.Department, error) {
    return s.storage.GetOrganizationTree()
}

func (s *serviceImpl) SearchEmployees(position, skill string, departmentID *int) ([]models.Employee, error) {
    return s.storage.SearchEmployees(position, skill, departmentID)
}

func (s *serviceImpl) GetEmployeeDetails(employeeID int) (*models.Employee, error) {
    return s.storage.GetEmployeeDetails(employeeID)
}

func (s *serviceImpl) UpdateEmployee(employee models.Employee) error {
    return s.storage.UpdateEmployee(employee)
}

func (s *serviceImpl) AddDepartment(department models.Department) error {
    return s.storage.AddDepartment(department)
}
