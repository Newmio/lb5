package models

type Course struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" binding:"required" db:"name"`
	TotalLessons int    `json:"total_lessons" db:"total_lessons"`
}
