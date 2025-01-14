package services

import (
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/repositories"
)

func ValidateUser(user models.User) bool {
	return true
}

func SaveUser(usr *models.User) error {
	err := repositories.SaveUser(usr)
	return err
}
