package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/oscarvo29/real-chat-backend/models"
)

func SaveMessage(message *models.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	uuid := uuid.New()

	query := `
	INSERT INTO messages (
		message_uuid,
		sender_uuid,
		receiver_uuid,
		message_value,
		send_time 
		) VALUES ($1, $2, $3, $4, $5)
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid, message.SenderUuid, message.ReceiverUuid, message.MessageValue, time.Now())
	if err != nil {
		return err
	}

	message.MessageId = uuid
	return nil
}

func ChatHistory(activeUuid, chatUuid uuid.UUID) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	fmt.Println("Chatuuid: ", chatUuid.String())

	var chatMessages []*models.Message

	query := `SELECT * FROM messages WHERE sender_uuid=$1 AND receiver_uuid=$2 OR sender_uuid=$2 AND receiver_uuid=$1`
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return chatMessages, err
	}

	row, err := stmt.Query(activeUuid, chatUuid)
	if err != nil {
		return chatMessages, err
	}

	for row.Next() {
		var msg models.Message
		err = row.Scan(
			&msg.MessageId,
			&msg.SenderUuid,
			&msg.ReceiverUuid,
			&msg.MessageValue,
			&msg.SendTime,
			&msg.Read,
			&msg.ReadAt,
		)
		if err != nil {
			return chatMessages, err
		}

		chatMessages = append(chatMessages, &msg)
	}

	return chatMessages, nil
}
