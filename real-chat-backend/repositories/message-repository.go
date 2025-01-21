package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
		message_id,
		chat_id,
		sender_id,
		message_value,
		send_time 
		) VALUES ($1, $2, $3, $4, $5)
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, uuid, message.ChatId, message.SenderUuid, message.MessageValue, time.Now())
	if err != nil {
		return err
	}

	message.MessageId = uuid
	return nil
}

func SaveChat(chat *models.Chat) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	uuid := uuid.New()

	query := `INSERT INTO chats (
		chat_id,
		chat_name
	) VALUES ($1, $2)`

	_, err := DB.ExecContext(ctx, query, uuid, chat.ChatName)
	if err != nil {
		return err
	}

	chat.ChatUuid = uuid
	return err
}

func RegisterChatParticipant(chatId, participantId uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO chat_participants (
		chat_id,
		uuid
	) VALUES ($1, $2)`

	_, err := DB.ExecContext(ctx, query, chatId, participantId)
	return err
}

func GetAllChatsForUser(activeUuid uuid.UUID) ([]*models.Chat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var chats []*models.Chat

	query := `
		SELECT chats.chat_id, latest_message_id, chat_name
		FROM chats
		INNER JOIN chat_participants AS cp on chats.chat_id = cp.chat_id
		WHERE cp.uuid = $1;
	`

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return chats, err
	}

	row, err := stmt.Query(activeUuid)
	if err != nil {
		return chats, err
	}

	for row.Next() {
		var chat models.Chat
		var latestMessageId sql.NullString

		err = row.Scan(
			&chat.ChatUuid,
			&latestMessageId,
			&chat.ChatName,
		)

		if err != nil {
			return chats, err
		}

		if latestMessageId.Valid {
			msg, err := GetMessage(latestMessageId.String)
			if err != nil {
				log.Fatal(err)
			}
			chat.LatestMessage = msg
		} else {
			chat.LatestMessage = nil
		}
		chats = append(chats, &chat)
	}

	return chats, nil
}

func GetChatHistory(chatId uuid.UUID, senderUuid uuid.UUID) ([]*models.Message, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		SELECT * FROM messages
		WHERE chat_id =  $1;
	`
	var messages []*models.Message

	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return messages, err
	}

	row, err := stmt.Query(chatId)
	if err != nil {
		return messages, err
	}

	for row.Next() {
		var msg models.Message

		err = row.Scan(
			&msg.MessageId,
			&msg.ChatId,
			&msg.SenderUuid,
			&msg.MessageValue,
			&msg.SendTime,
			&msg.Read,
			&msg.ReadAt,
		)

		if err != nil {
			return messages, err
		}

		msg.IsSender = msg.SenderUuid == senderUuid
		messages = append(messages, &msg)
	}

	return messages, nil
}

// func GetMessage(messageId string, msg *models.Message) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
// 	defer cancel()
// 	fmt.Println("Message ID: ", messageId)

// 	query := `SELECT * FROM messages WHERE message_id = $1;`
// 	row, err := DB.QueryContext(ctx, query, messageId)
// 	if err != nil {
// 		return err
// 	}

// 	err = row.Scan(
// 		&msg.MessageId,
// 		&msg.ChatId,
// 		&msg.SenderUuid,
// 		&msg.MessageValue,
// 		&msg.SendTime,
// 		&msg.Read,
// 		&msg.ReadAt,
// 	)

// 	return err
// }

func GetMessage(messageId string) (*models.Message, error) {
	var msg models.Message
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	fmt.Println("Message ID: ", messageId)

	query := `SELECT * FROM messages WHERE message_id = $1;`
	row, err := DB.QueryContext(ctx, query, messageId)
	if err != nil {
		return &msg, err
	}

	for row.Next() {
		err = row.Scan(
			&msg.MessageId,
			&msg.ChatId,
			&msg.SenderUuid,
			&msg.MessageValue,
			&msg.SendTime,
			&msg.Read,
			&msg.ReadAt,
		)
	}

	return &msg, err
}

func GetReceipiants(chatId uuid.UUID) ([]*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	var receipiants []*string

	query := `
		SELECT uuid FROM chat_participants WHERE chat_id = $1;
	`
	stmt, err := DB.PrepareContext(ctx, query)
	if err != nil {
		return receipiants, err
	}

	row, err := stmt.Query(chatId)
	if err != nil {
		return receipiants, err
	}

	for row.Next() {
		var receipiant uuid.UUID
		err = row.Scan(&receipiant)
		if err != nil {
			return receipiants, err
		}
		idString := receipiant.String()
		receipiants = append(receipiants, &idString)
	}

	return receipiants, nil
}
