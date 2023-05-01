package student

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

func GroupById(db *sqlx.DB) ([]models.StudentsAge, error) {
	var students []models.StudentsAge
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select name, sum(age) as all_age from students group by name"

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
