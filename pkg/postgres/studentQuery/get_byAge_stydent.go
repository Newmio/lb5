package student

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

func GetByAge(db *sqlx.DB, min, max int) ([]models.Student, error) {
	var students []models.Student
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select * from students where age>=$1 and age<=$2"

	err = tx.Select(&students, str, min, max)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return students, nil
}
