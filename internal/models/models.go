package models

type Employee struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Position    string   `json:"position"`
    Skills      []string `json:"skills"`
    DepartmentID int     `json:"department_id"`
}

type Department struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    ParentID *int   `json:"parent_id,omitempty"`
}