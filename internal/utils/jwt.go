package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

var (
	SecretKey     = []byte(os.Getenv("JWT_SECRET"))
	Expiration, _ = strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))
)

func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["exp"] = time.Now().Add(time.Duration(Expiration)).Unix()

	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
