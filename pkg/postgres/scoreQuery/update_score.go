package scoreQuery

import (
	"errors"
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// обновляет оценку из указанного курса и указанного студента
func Update(db *sqlx.DB, studentId, courseId, scoreId int, input models.Score) error {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	str := "update scores set score_value=$1 where id=$2 and id_course=$3 and id_student=$4"

	result, err := tx.Exec(str, input.ScoreValue, scoreId, courseId, scoreId)
	if err != nil {
		tx.Rollback()
		return err
	}

	row, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if row == 0 {
		tx.Rollback()
		return errors.New("error in update_score.go, line 31")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
