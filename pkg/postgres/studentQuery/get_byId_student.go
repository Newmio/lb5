package student

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// выводит по айди студента у указанного юзера
func GetById(db *sqlx.DB, userId, studentId int) (models.Student, error) {
	var student models.Student
	tx, err := db.Beginx()
	if err != nil {
		return student, err
	}

	str := "select students.* from students join users_students on students.id = users_students.id_student where users_students.id_user = $1 and users_students.id_student = $2"

	err = tx.Get(&student, str, userId, studentId)
	if err != nil {
		tx.Rollback()
		return student, err
	}

	return student, err
}
