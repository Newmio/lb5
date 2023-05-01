package scoreQuery

import (
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// выводит все оценки у указанного курса и указанного студента
func GetAll(db *sqlx.DB, studentId, courseId int) ([]models.Score, error) {
	var scores []models.Score
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	str := "select * from scores where id_course = $1 and id_student = $2"

	err = tx.Select(&scores, str, courseId, studentId)
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
