package storage

import (
	"context"
	"encoding/json"
    "fmt"
	"strconv"
	"strings"
	"time"

	"WEB-DA/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewStorage(db *sqlx.DB, redis *redis.Client) *Storage {
	return &Storage{db: db, redis: redis}
}

// Получить всех сотрудников с использованием Redis
func (s *Storage) GetAllEmployees() ([]models.Employee, error) {
	ctx := context.Background()
	cacheKey := "employees:all"

	// Проверяем кэш
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var employees []models.Employee
		if err := json.Unmarshal([]byte(cached), &employees); err == nil {
			return employees, nil
		}
	}

	// Если кэш пуст, получаем данные из PostgreSQL
	query := `SELECT * FROM employees ORDER BY id`
	var employees []models.Employee
	if err := s.db.Select(&employees, query); err != nil {
		return nil, err
	}

	// Сохраняем результат в Redis с TTL = 60 секунд
	data, _ := json.Marshal(employees)
	_ = s.redis.Set(ctx, cacheKey, data, 60*time.Second).Err()

	return employees, nil
}

// Получить сотрудников с фильтрами и использованием Redis
func (s *Storage) GetEmployeesWithFilters(filters map[string]interface{}) ([]models.Employee, error) {
	ctx := context.Background()
	cacheKey := "employees:filters:" + s.generateFilterKey(filters)

	// Проверяем кэш
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var employees []models.Employee
		if err := json.Unmarshal([]byte(cached), &employees); err == nil {
			return employees, nil
		}
	}

	// Формируем SQL-запрос с фильтрами
	query := "SELECT * FROM employees"
	var conditions []string
	var args []interface{}

	for key, value := range filters {
		if value != "" && value != nil {
			conditions = append(conditions, key+" = $"+strconv.Itoa(len(args)+1))
			args = append(args, value)
		}
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// Выполняем запрос
	var employees []models.Employee
	if err := s.db.Select(&employees, query, args...); err != nil {
		return nil, err
	}

	// Сохраняем результат в Redis с TTL = 60 секунд
	data, _ := json.Marshal(employees)
	_ = s.redis.Set(ctx, cacheKey, data, 60*time.Second).Err()

	return employees, nil
}

// Получить иерархию сотрудников с использованием Redis
func (s *Storage) GetEmployeeHierarchy(managerID interface{}, departmentID interface{}) ([]models.Employee, error) {
	ctx := context.Background()
	cacheKey := "employees:hierarchy:" + s.generateHierarchyKey(managerID, departmentID)

	// Проверяем кэш
	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var employees []models.Employee
		if err := json.Unmarshal([]byte(cached), &employees); err == nil {
			return employees, nil
		}
	}

	// Если кэш пуст, выполняем SQL-запрос
	query := `
        WITH RECURSIVE employee_hierarchy AS (
            SELECT DISTINCT 
                e.id, e.first_name, e.last_name, e.position, e.role, e.phone, e.email,
                e.manager_id, e.department_id, e.division_id, e.unit_id
            FROM employees e
            WHERE ($1::INT IS NULL OR e.manager_id = $1)
              AND ($2::INT IS NULL OR e.department_id = $2)

            UNION ALL

            SELECT DISTINCT 
                e.id, e.first_name, e.last_name, e.position, e.role, e.phone, e.email,
                e.manager_id, e.department_id, e.division_id, e.unit_id
            FROM employees e
            INNER JOIN employee_hierarchy eh ON e.manager_id = eh.id
        )
        SELECT DISTINCT * 
        FROM employee_hierarchy
        ORDER BY manager_id NULLS FIRST, id;
    `
	var employees []models.Employee
	if err := s.db.Select(&employees, query, managerID, departmentID); err != nil {
		return nil, err
	}

	// Сохраняем результат в Redis с TTL = 60 секунд
	data, _ := json.Marshal(employees)
	_ = s.redis.Set(ctx, cacheKey, data, 60*time.Second).Err()

	return employees, nil
}

// Вспомогательный метод для генерации ключа кэша по фильтрам
func (s *Storage) generateFilterKey(filters map[string]interface{}) string {
	var filterParts []string
	for key, value := range filters {
		filterParts = append(filterParts, key+" LIKE $"+fmt.Sprintf("%v", value))
	}
	return strings.Join(filterParts, "&")
}

// Вспомогательный метод для генерации ключа кэша для иерархии
func (s *Storage) generateHierarchyKey(managerID interface{}, departmentID interface{}) string {
	return fmt.Sprintf("manager_id=%v&department_id=%v", managerID, departmentID)
}

// GetFilters возвращает список доступных фильтров
func (s *Storage) GetFilters() ([]models.Filters, error) {
    filters := []models.Filters{
        {Key: "first_name", Name: "Имя сотрудника"},
        {Key: "last_name", Name: "Фамилия сотрудника"},
        {Key: "position", Name: "Должность сотрудника"},
        {Key: "role", Name: "Роль сотрудника"},
        {Key: "phone", Name: "Номер телефона"},
        {Key: "email", Name: "Электронная почта"},
        {Key: "manager_id", Name: "ID руководителя"},
        {Key: "department_id", Name: "ID департамента"},
        {Key: "division_id", Name: "ID подразделения"},
        {Key: "unit_id", Name: "ID отдела"},
    }
    return filters, nil
}

// GetSectionName возвращает название отдела по ID
func (s *Storage) GetSectionName(id int) (string, error) {
    var name string
    err := s.db.Get(&name, "SELECT name FROM departments WHERE id = $1", id)
    if err != nil {
        return "", err
    }
    return name, nil
}