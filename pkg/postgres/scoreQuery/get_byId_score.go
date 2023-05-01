package scoreQuery

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// выводит по айди оценку из указанного курса и указанного студента
func GetById(db *sqlx.DB, courseId, studentId, scoreId int) (models.Score, error) {
	var score models.Score
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return score, err
	}

	str := "select * from scores where id = $1 and id_student = $2 and id_course = $3"

	err = tx.Get(&score, str, scoreId, studentId, courseId)
	if err != nil {
		tx.Rollback()
		return score, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return score, err
	}

	return score, nil
}
