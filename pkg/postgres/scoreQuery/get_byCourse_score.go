package scoreQuery

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

func GetByCourse(db *sqlx.DB, courseId int) ([]models.Score, error) {
	var scores []models.Score
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select * from scores where id_course=$1"

	err = tx.Select(&scores, str, courseId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}

	return scores, nil
}
