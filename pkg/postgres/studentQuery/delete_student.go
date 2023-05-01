package student

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// удаляет студента по айди у указанного юзера
func Delete(db sqlx.DB, userId, studentId int) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	str := "delete from students where id = $1 and id in(select id_student from users_students where id_user = $2)"

	result, err := tx.Exec(str, studentId, userId)
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
		return errors.New("error in delete_course.go, line 29")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
