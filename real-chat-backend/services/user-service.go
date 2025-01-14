package services

import (
	"errors"

	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/repositories"
	"golang.org/x/crypto/bcrypt"
)

func ValidateUser(clientUser *models.User) (bool, error) {
	databaseUser, err := repositories.GetUserFromName(clientUser.Name)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(clientUser.Password))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil // Password is wrong.
		default:
			return false, err // An error accured.
		}
	}

	clientUser.Uuid = databaseUser.Uuid
	return true, nil
}

func SaveUser(usr *models.User) error {
	err := repositories.SaveUser(usr)
	return err
}

func GetAllUsers() ([]*models.User, error) {
	return repositories.GetAllUsers()
}
