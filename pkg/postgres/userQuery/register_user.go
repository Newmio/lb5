package userquery

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"lb5/pkg/models"

	"github.com/jmoiron/sqlx"
)

// регистрация юзера
func Post(db *sqlx.DB, user models.User) error {
	tx, err := db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	user.Password = generateHash(user.Password)

	str := "insert into users(login, password) values($1, $2)"

	result, err := tx.Exec(str, user.Login, user.Password)
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
		return errors.New("error in register_user.go, line 33")
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func generateHash(pass string) string {
	hash := sha1.New()
	hash.Write([]byte(pass))
	return fmt.Sprintf("%x", hash.Sum([]byte("dasdasda")))
}
