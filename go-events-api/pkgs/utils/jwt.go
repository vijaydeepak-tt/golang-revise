package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret-key"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse the token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid token claims")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil

}
