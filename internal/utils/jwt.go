package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"strconv"
	"time"
)

func GenerateToken(username string) (string, error) {
	expirationTime, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Duration(expirationTime) * time.Minute).Unix()

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
