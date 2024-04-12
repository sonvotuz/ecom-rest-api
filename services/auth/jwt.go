package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/vnsonvo/ecom-rest-api/types"
	"github.com/vnsonvo/ecom-rest-api/utils"
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

func JWTAuthMiddleware(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the token from request
		tokenString := r.Header.Get("Authorization")

		// validate JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWTSECRET")), nil
		})
		if err != nil {
			log.Printf("Failed to validate token: %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Println("Invalid token")
			permissionDenied(w)
			return
		}

		// fetch user from DB using userId from claims
		claims := token.Claims.(jwt.MapClaims)
		str := claims["userId"].(string)

		userId, err := strconv.Atoi(str)
		if err != nil {
			log.Println("Invalid user id")
			permissionDenied(w)
			return
		}

		// verify user id is valid
		user, err := store.GetUserByID(userId)
		if err != nil {
			log.Printf("Failed to get user by id: %v", err)
			permissionDenied(w)
			return
		}

		// set context "userId" to user id
		ctx := r.Context()
		ctx = context.WithValue(ctx, "userId", user.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("Permission denied"))
}

func GetUserIDFromContext(ctx context.Context) int {
	userId, ok := ctx.Value("userId").(int)
	if !ok {
		return -1
	}
	return userId
}
