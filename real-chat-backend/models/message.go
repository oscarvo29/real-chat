package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	MessageId    uuid.UUID    `json:"message_id,omitempty"`
	SenderUuid   uuid.UUID    `json:"sender_uuid,omitempty"`
	ReceiverUuid uuid.UUID    `json:"receiver_uuid"`
	MessageValue string       `json:"message_value"`
	SendTime     time.Time    `json:"send_time,omitempty"`
	Read         bool         `json:"read,omitempty"`
	ReadAt       sql.NullTime `json:"read_at,omitempty"`
}

type ShortMessage struct {
	SenderJWT    string    `json:"jwt"`
	ReceiverUuid uuid.UUID `json:"receiver_uuid"`
	Message      string    `json:"message"`
}

func NewMessage(senderUiid, receiverUuid uuid.UUID, messageValue string) *Message {
	return &Message{
		SenderUuid:   senderUiid,
		ReceiverUuid: receiverUuid,
		MessageValue: messageValue,
		SendTime:     time.Now(),
		Read:         false,
	}
}
