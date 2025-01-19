package models

import "github.com/google/uuid"

type ChatParticipantsIds []string

type Chat struct {
	ChatUuid      uuid.UUID           `json:"chat_uuid"`
	Participants  ChatParticipantsIds `json:"participants,omitempty"`
	ChatName      string              `json:"chat_name"`
	LatestMessage *Message            `json:"latest_message,omitempty"`
	Messages      []Message           `json:"messages,omitempty"`
}

type ChatParticipants struct {
	ChatUuid uuid.UUID `json:"chat_uuid"`
	UserUuid uuid.UUID `json:"participant_uuid`
}
