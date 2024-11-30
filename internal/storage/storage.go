package storage

import (
	"WEB-DA/internal/models"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// Получить иерархию сотрудников с использованием рекурсии
func (s *Storage) GetEmployeeHierarchy(managerID interface{}, departmentID interface{}) ([]models.Employee, error) {
    query := `
        WITH RECURSIVE employee_hierarchy AS (
            SELECT 
                e.id, e.first_name, e.last_name, e.position, e.role, e.phone, e.email,
                e.manager_id, e.department_id, e.division_id, e.unit_id
            FROM employees e
            WHERE ($1::INT IS NULL OR e.manager_id = $1)
              AND ($2::INT IS NULL OR e.department_id = $2)

            UNION ALL

            SELECT 
                e.id, e.first_name, e.last_name, e.position, e.role, e.phone, e.email,
                e.manager_id, e.department_id, e.division_id, e.unit_id
            FROM employees e
            INNER JOIN employee_hierarchy eh ON e.manager_id = eh.id
        )
        SELECT * FROM employee_hierarchy;
    `

    var employees []models.Employee
    err := s.db.Select(&employees, query, managerID, departmentID)
    return employees, err
}

// Получить всех сотрудников
func (s *Storage) GetAllEmployees() ([]models.Employee, error) {
	query := `SELECT * FROM employees`
	var employees []models.Employee
	err := s.db.Select(&employees, query)
	return employees, err
}
