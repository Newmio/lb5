package student

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

func OrderBy(db *sqlx.DB) ([]models.Student, error) {
	var students []models.Student
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select * from students order by age"

	err = tx.Select(&students, str)
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
