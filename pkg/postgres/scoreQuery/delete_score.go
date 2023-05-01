package scoreQuery

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// удаляет оценку из указанного курса и указанного студента
func Delete(db *sqlx.DB, courseId, studentId, scoreId int) error {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	str := "delete from scores where id = $1 and id_student = $2 and id_course = $3"

	result, err := tx.Exec(str, scoreId, studentId, courseId)
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
		return errors.New("error in delete_score.go, line 30")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
