package services

import (
	"github.com/google/uuid"
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/repositories"
)

func SaveMessage(msg *models.Message) error {
	return repositories.SaveMessage(msg)
}

func GetAllChatMessages(activeUuid, chatUuid uuid.UUID) ([]*models.Message, error) {
	return repositories.ChatHistory(activeUuid, chatUuid)
}

