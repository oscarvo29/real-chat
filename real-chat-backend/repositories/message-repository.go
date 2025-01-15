package repositories

import (
	"context"
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
