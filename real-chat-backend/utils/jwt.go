package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(name string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"uuid": userId,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(JWT_SECRET))
}

func VerifyToken(tokenInput string) (int64, error) {
	parsedToken, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userId := int64(claims["uuid"].(float64))

	return userId, nil
}
