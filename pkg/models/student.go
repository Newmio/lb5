package models

type Student struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	LastName string `json:"last_name" binding:"required" db:"last_name"`
	Age      int    `json:"age" binding:"required" db:"age"`
}
