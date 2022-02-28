package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

func GenerateToken(username string, secret string, expiration string) (string, error) {
	expirationTime, _ := strconv.Atoi(expiration)
	secretKey := []byte(secret)
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
