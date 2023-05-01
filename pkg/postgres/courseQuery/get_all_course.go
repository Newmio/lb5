package course

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// выводит все курсы, которые связаня с Х студентом
func GetAll(db *sqlx.DB, studentId int) ([]models.Course, error) {
	var courses []models.Course
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select courses.* from courses join students_courses on courses.id = students_courses.id_course where students_courses.id_student=$1"

	err = tx.Select(&courses, str, studentId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return courses, nil
}
