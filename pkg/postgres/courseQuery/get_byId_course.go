package course

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// выводит курс по айди у указанного студента
func GetById(db *sqlx.DB, studentId, courseId int) (models.Course, error) {
	var course models.Course

	tx, err := db.Beginx()
	if err != nil {
		return course, err
	}

	str := "select courses.* from courses join students_courses on courses.id = students_courses.id_course where students_courses.id_course=$1 and students_courses.id_student=$2"

	err = tx.Get(&course, str, courseId, studentId)
	if err != nil {
		return course, err
	}

	if err := tx.Commit(); err != nil {
		return course, err
	}

	return course, nil
}
