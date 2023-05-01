package course

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// удаляет курс у указанного студента
func Delete(db *sqlx.DB, studentId, courseId int) error {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	str := "delete from courses where id=$1 and id in(select id_student from students_courses where students_courses.id_student=$2)"

	result, err := tx.Exec(str, courseId, studentId)
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
		return errors.New("error in delete_course.go, line 27")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
