package services

import (
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/repositories"
)

func SaveMessage(msg *models.Message) error {
	return repositories.SaveMessage(msg)
}
