package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(name string, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"uuid": uuid,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(JWT_SECRET))
}

func VerifyToken(tokenInput string) (string, error) {
	parsedToken, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		fmt.Println(err)
		return "", errors.New("could not parse token")
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid token claims")
	}

	uuid, ok := claims["uuid"].(string)
	if !ok {
		return "", errors.New("failed to convert the uuid when verifying user")
	}

	return uuid, nil
}
