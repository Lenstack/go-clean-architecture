package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

var (
	SecretKey = []byte(os.Getenv("JWT_SECRET"))
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
