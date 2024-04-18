package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(UserID int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = UserID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	return token.SignedString([]byte("secret-key"))
}
