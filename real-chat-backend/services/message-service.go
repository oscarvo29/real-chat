package services

import (
	"github.com/google/uuid"
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/repositories"
)

func SaveMessage(msg *models.Message) error {
	return repositories.SaveMessage(msg)
}

func GetChatHistory(chatId string, senderUuid uuid.UUID) ([]*models.Message, error) {
	chatUuid, err := uuid.Parse(chatId)
	if err != nil {
		return []*models.Message{}, err
	}

	return repositories.GetChatHistory(chatUuid, senderUuid)
}

func GetChatRooms(activeUuid uuid.UUID) ([]*models.Chat, error) {
	return repositories.GetAllChatsForUser(activeUuid)
}

func CreateChat(chat *models.Chat) error {
	err := repositories.SaveChat(chat)
	if err != nil {
		return err
	}

	for _, participant := range chat.Participants {
		uuid, err := uuid.Parse(participant)
		if err != nil {
			return err
		}

		err = repositories.RegisterChatParticipant(chat.ChatUuid, uuid)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetReceipiants(chatId uuid.UUID) ([]*string, error) {
	return repositories.GetReceipiants(chatId)
}
