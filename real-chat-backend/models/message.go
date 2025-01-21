package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageId    uuid.UUID     `json:"message_id,omitempty"`
	ChatId       uuid.UUID     `json:"chat_id"`
	SenderUuid   uuid.UUID     `json:"sender_uuid,omitempty"`
	MessageValue string       `json:"message_value"`
	SendTime     time.Time    `json:"send_time,omitempty"`
	Read         bool         `json:"read,omitempty"`
	ReadAt       sql.NullTime `json:"read_at,omitempty"`
	IsSender     bool          `json:"is_sender"`
}

type ShortMessage struct {
	SenderJWT string    `json:"jwt"`
	ChatId    uuid.UUID `json:"chat_id"`
	Message   string    `json:"message"`
}

func NewMessage(senderUiid, chatId uuid.UUID, messageValue string) *Message {
	now := time.Now()
	read := false

	return &Message{
		SenderUuid:   senderUiid,
		ChatId:       chatId,
		MessageValue: messageValue,
		SendTime:     now,
		Read:         read,
	}
}
