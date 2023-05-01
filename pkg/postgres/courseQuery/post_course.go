package course

import (
	"errors"
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// создает курс у указанного студента
func Post(db *sqlx.DB, studentId int, course models.Course) (int, error) {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	str := "insert into courses(name, total_lessons) values($1, $2) returning id"

	var id int
	row := tx.QueryRow(str, course.Name, course.TotalLessons)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	str = "insert into students_courses(id_student, id_course) values($1, $2)"

	result, err := tx.Exec(str, studentId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	if rows == 0 {
		tx.Rollback()
		return 0, errors.New("error in post_course.go, line 41")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}
