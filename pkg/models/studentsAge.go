package models

type StudentsAge struct {
	Name   string `json:"name" db:"name"`
	AllAge int    `json:"all_age" db:"all_age"`
}
