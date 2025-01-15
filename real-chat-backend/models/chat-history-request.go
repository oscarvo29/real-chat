package models

import (
	"github.com/google/uuid"
)

// type ChatHistory struct {
// 	ReceiverUuidStr string    `json:"receiver_uuid"`
// 	ReceiverUuid    uuid.UUID `json:"-"`
// }

// func (c *ChatHistory) SetUuid() {
// 	uuidObj, err := uuid.Parse(c.ReceiverUuidStr)
// 	if err != nil {
// 		fmt.Println("Error happend when trying to parse uuid.")
// 		return
// 	}

// 	c.ReceiverUuid = uuidObj
// }

type ChatHistory struct {
	ReceiverUuid uuid.UUID `json:"receiver_uuid"`
}
