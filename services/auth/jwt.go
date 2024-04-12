package auth

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	expireTimeStr := os.Getenv("JWTEXPIREINSECONDS")
	if expireTimeStr == "" {
		expireTimeStr = "0"
	}
	expireTime, err := strconv.Atoi(expireTimeStr)
	if err != nil {
		return "", err
	}

	expiration := time.Second * time.Duration(expireTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
