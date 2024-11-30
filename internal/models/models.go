package models

type Employee struct {
    ID           int     `json:"id" db:"id"`  // db:"id" указывает на столбец id в базе данных
    FirstName    string  `json:"first_name" db:"first_name"`
    LastName     string  `json:"last_name" db:"last_name"`
    Position     string  `json:"position" db:"position"`
    Role         string  `json:"role" db:"role"`
    Phone        string  `json:"phone" db:"phone"`
    Email        string  `json:"email" db:"email"`
    ManagerID    *int    `json:"manager_id" db:"manager_id"`
    DepartmentID *int    `json:"department_id" db:"department_id"`
    DivisionID   *int    `json:"division_id" db:"division_id"`
    UnitID       *int    `json:"unit_id" db:"unit_id"`
}

type Department struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	HeadID *int    `json:"head_id"`
}

type Division struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	HeadID *int    `json:"head_id"`
}

type Unit struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	HeadID *int    `json:"head_id"`
}

type Filters struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}
