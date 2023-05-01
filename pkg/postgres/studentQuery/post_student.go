package student

import (
	"errors"
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// создает студента у указанного юзера
func Post(db *sqlx.DB, student models.Student, userId int) (int, error) {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	str := "insert into students(name, last_name, age) values($1, $2, $3) returning id"
	row := tx.QueryRow(str, student.Name, student.LastName, student.Age)
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	str = "insert into users_students(id_user, id_student) values($1, $2)"
	result, err := tx.Exec(str, userId, id)
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
		return 0, errors.New("error in post_student.go, line 41")
	}
	return id, tx.Commit()
}
