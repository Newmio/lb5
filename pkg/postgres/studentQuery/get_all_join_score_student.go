package student

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

func GetAllJoinScore(db *sqlx.DB) ([]models.StudentScore, error) {
	var studentScore []models.StudentScore
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select students.id, students.name, students.last_name, students.age, scores.id, scores.score_value, scores.id_course, scores.id_student from students join scores on students.id = scores.id_student"

	rows, err := tx.Queryx(str)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	for rows.Next() {
		var student models.Student
		var score models.Score

		err = rows.Scan(&student.Id, &student.Name, &student.LastName, &student.Age, &score.Id, &score.ScoreValue,
			&score.IdCourse, &score.IdStudent)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		studentScore = append(studentScore, models.StudentScore{Student: student, Score: score})
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return studentScore, nil
}
