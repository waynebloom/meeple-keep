package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "superSecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":  email,
			"userId": userId,
			"exp":    time.Now().Add(time.Hour * 2).Unix(),
		})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenStr string) (int64, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Token was signed with an unexpected signing method.")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("An invalid token was provided.")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Token contained claims with an unexpected type.")
	}

	// email := claims["email"].(string)
	id := claims["userId"].(float64)
	// exp := claims["exp"].(int64)

	return int64(id), nil
}
