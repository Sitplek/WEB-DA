package models

import "time"

type Department struct {
	ID             uint         `gorm:"primaryKey" json:"ID"`
	Name           string       `json:"Name"`
	ParentID       *uint        `json:"ParentID"`
	Parent         *Department  `gorm:"foreignKey:ParentID" json:"Parent,omitempty"`
	SubDepartments []Department `gorm:"foreignKey:ParentID" json:"SubDepartments,omitempty"`
	CreatedAt      time.Time    `json:"CreatedAt"`
	UpdatedAt      time.Time    `json:"UpdatedAt"`
	Employees      []Employee   `json:"Employees,omitempty"`
}

type Employee struct {
	ID           uint       `gorm:"primaryKey" json:"ID"`
	Name         string     `json:"Name"`
	Email        string     `json:"Email"`
	Phone        string     `json:"Phone"`
	Address      string     `json:"Address"`
	Role         string     `json:"Role"`
	JobPosition  string     `json:"JobPosition"`
	SupervisorID *uint      `json:"SupervisorID"`
	Supervisor   *Employee  `gorm:"foreignKey:SupervisorID" json:"Supervisor,omitempty"`
	DepartmentID uint       `json:"DepartmentID"`
	Department   Department `gorm:"foreignKey:DepartmentID" json:"Department"`
	CreatedAt    time.Time  `json:"CreatedAt"`
	UpdatedAt    time.Time  `json:"UpdatedAt"`
}
