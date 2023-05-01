package student

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// выводит всех студентов у указанного юзера
func GetAll(db *sqlx.DB, userId int) ([]models.Student, error) {
	var students []models.Student
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}

	str := "select students.* from students join users_students on students.id = users_students.id_student where users_students.id_user = $1 "

	err = tx.Select(&students, str, userId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return students, err
}
