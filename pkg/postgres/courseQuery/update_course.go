package course

import (
	"errors"
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// обновляет курс у указанного студента
func Update(db *sqlx.DB, studentId, courseId int, input models.Course) error {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	str := "update courses set name=$1, total_lessons=$2 from students_courses where courses.id = students_courses.id_course and students_courses.id_student=$3 and courses.id=$4"

	result, err := tx.Exec(str, input.Name, input.TotalLessons, studentId, courseId)
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
		return errors.New("error in update_course.go, line 28")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
