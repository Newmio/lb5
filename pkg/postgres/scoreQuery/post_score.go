package scoreQuery

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// создает оценку у указанного курса и указанного студента
func Post(db *sqlx.DB, input models.Score, studentId, courseId int) (int, error) {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	str := "insert into scores(score_value, id_course, id_student) values($1, $2, $3) returning id"

	row := tx.QueryRow(str, input.ScoreValue, courseId, studentId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		tx.Commit()
		return 0, err
	}

	return id, nil
}
