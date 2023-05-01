package models

// промежуточная структура для соеденения юзера и его всех оценок
type StudentScore struct {
	Student Student `json:"student"`
	Score   Score   `json:"score"`
}
