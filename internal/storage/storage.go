package storage

import (
    "database/sql"
    "errors"
    "WEB-DA/internal/models"
	"strings"
)

type Storage interface {
    GetOrganizationTree() ([]models.Department, error)
    SearchEmployees(position, skill string, departmentID *int) ([]models.Employee, error)
    GetEmployeeDetails(employeeID int) (*models.Employee, error)
    UpdateEmployee(employee models.Employee) error
    AddDepartment(department models.Department) error
}

type storageImpl struct {
    db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
    return &storageImpl{db: db}
}

func (s *storageImpl) GetOrganizationTree() ([]models.Department, error) {
    rows, err := s.db.Query(`SELECT id, name, parent_id FROM departments`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var departments []models.Department
    for rows.Next() {
        var dep models.Department
        err := rows.Scan(&dep.ID, &dep.Name, &dep.ParentID)
        if err != nil {
            return nil, err
        }
        departments = append(departments, dep)
    }
    return departments, nil
}

func (s *storageImpl) SearchEmployees(position, skill string, departmentID *int) ([]models.Employee, error) {
    query := `SELECT id, name, position, skills, department_id FROM employees WHERE position ILIKE $1`
    args := []interface{}{"%" + position + "%"}
    if departmentID != nil {
        query += " AND department_id = $2"
        args = append(args, *departmentID)
    }
    rows, err := s.db.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var employees []models.Employee
    for rows.Next() {
        var emp models.Employee
        var skills string
        if err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &skills, &emp.DepartmentID); err != nil {
            return nil, err
        }
        emp.Skills = strings.Split(skills, ",")
        employees = append(employees, emp)
    }
    return employees, nil
}

func (s *storageImpl) GetEmployeeDetails(employeeID int) (*models.Employee, error) {
    row := s.db.QueryRow(`SELECT id, name, position, skills, department_id FROM employees WHERE id = $1`, employeeID)
    var emp models.Employee
    var skills string
    if err := row.Scan(&emp.ID, &emp.Name, &emp.Position, &skills, &emp.DepartmentID); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, nil
        }
        return nil, err
    }
    emp.Skills = strings.Split(skills, ",")
    return &emp, nil
}

func (s *storageImpl) UpdateEmployee(employee models.Employee) error {
    _, err := s.db.Exec(`UPDATE employees SET name = $1, position = $2, skills = $3, department_id = $4 WHERE id = $5`,
        employee.Name, employee.Position, strings.Join(employee.Skills, ","), employee.DepartmentID, employee.ID)
    return err
}

func (s *storageImpl) AddDepartment(department models.Department) error {
    _, err := s.db.Exec(`INSERT INTO departments (name, parent_id) VALUES ($1, $2)`, department.Name, department.ParentID)
    return err
}
