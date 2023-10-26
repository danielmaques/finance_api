package authentication

import (
	"time"

	"github.com/danielmaques/finance_api/config"
	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}
