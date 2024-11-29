package service

import (
    "WEB-DA/internal/models"
    "WEB-DA/internal/storage"
)

type OrganizationService struct {
    storage *storage.OrganizationStorage
}

func NewOrganizationService(storage *storage.OrganizationStorage) *OrganizationService {
    return &OrganizationService{storage: storage}
}

// Получение дерева организации
func (s *OrganizationService) GetOrganizationTree() ([]models.Department, error) {
    departments, err := s.storage.GetAllDepartments()
    if err != nil {
        return nil, err
    }

    var buildTree func([]models.Department, *uint) []models.Department
    buildTree = func(depts []models.Department, parentID *uint) []models.Department {
        var nodes []models.Department
        for _, d := range depts {
            if (parentID == nil && d.ParentID == nil) || (parentID != nil && d.ParentID != nil && *parentID == *d.ParentID) {
                // Для каждого департамента, добавляем его сотрудников
                d.Employees = s.getEmployeesForDepartment(d.ID, departments)
                d.SubDepartments = buildTree(depts, &d.ID)
                nodes = append(nodes, d)
            }
        }
        return nodes
    }

    return buildTree(departments, nil), nil
}

func (s *OrganizationService) getEmployeesForDepartment(departmentID uint, allDepartments []models.Department) []models.Employee {
    var employees []models.Employee
    for _, dept := range allDepartments {
        for _, emp := range dept.Employees {
            if emp.DepartmentID == departmentID {
                employees = append(employees, emp)
            }
        }
    }
    return employees
}
// Получение отфильтрованных сотрудников
func (s *OrganizationService) GetFilteredEmployees(filters map[string]interface{}) ([]models.Employee, error) {
    return s.storage.GetFilteredEmployees(filters)
}
