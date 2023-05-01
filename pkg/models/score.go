package models

type Score struct {
	Id         int `json:"id" db:"id"`
	ScoreValue int `json:"score_value" binding:"required" db:"score_value"`
	IdCourse   int `json:"id_course" db:"id_course"`
	IdStudent  int `json:"id_student" db:"id_student"`
}
