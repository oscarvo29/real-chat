package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ContextKey string

const UuidKey ContextKey = "uuid"

func GenerateToken(name string, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"uuid": uuid,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	JWT_KEY := GetEnvValue("JWT_TOKEN")
	return token.SignedString([]byte(JWT_KEY))
}

func VerifyToken(tokenInput string) (string, error) {
	parsedToken, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		JWT_KEY := GetEnvValue("JWT_TOKEN")
		return []byte(JWT_KEY ), nil
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

func TransFormJWT(r *http.Request) (uuid.UUID, error) {
	var senderUuid uuid.UUID

	senderUuidString, ok := r.Context().Value(UuidKey).(string)
	if !ok {
		return senderUuid, errors.New("error when trying to convert jwt to string")
	}

	if senderUuidString == "" {
		return senderUuid, errors.New("no jwt was found")
	}

	return uuid.Parse(senderUuidString)
}
