package userquery

import (
	"errors"
	"lb5/pkg/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
)

// логин юзера
func Auth(db *sqlx.DB, login, pass string) (models.User, error) {
	var user models.User
	tx, err := db.Beginx()
	if err != nil {
		return user, err
	}

	strCommand := "select id from users where login=$1 and password=$2"

	err = tx.Get(&user, strCommand, login, pass)
	if err != nil {
		return user, err
	}
	return user, nil
}

type claims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GenerateToken(db *sqlx.DB, login, pass string) (string, error) {
	user, err := Auth(db, login, generateHash(pass))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id})

	return token.SignedString([]byte("dasdaasdasd"))
}

func ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte("dasdaasdasd"), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*claims)
	if !ok {
		return 0, errors.New("token claims are not of type *claims")
	}

	return claims.UserId, nil
}
