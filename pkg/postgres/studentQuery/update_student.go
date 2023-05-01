package student

import (
	"errors"
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// обновлет студента по айди у указанного юзера
func Update(db sqlx.DB, userId, studentId int, input models.Student) error {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}

	str := "update students set name=$1, last_name=$2, age=$3 from users_students where students.id = users_students.id_student and users_students.id_user = $4 and students.id = $5"

	result, err := tx.Exec(str, input.Name, input.LastName, input.Age, userId, studentId)
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
		return errors.New("error in update_student.go, line 31")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
