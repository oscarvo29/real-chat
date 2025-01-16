package controllers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/services"
	"github.com/oscarvo29/real-chat-backend/utils"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Web socket have been hit?")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error when upgrading to websocket: ", err)
		return
	}
	defer conn.Close()

	fmt.Println("New WebSocket connection is openned.")

	for {
		var msg models.ShortMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("Error when reading message: ", err)
			return
		}

		uuidString, err := utils.VerifyToken(msg.SenderJWT)
		if err != nil {
			fmt.Println("Error when reading JWT: ", err)
			return
		}
		uuid, err := uuid.Parse(uuidString)
		if err != nil {
			fmt.Println("Error when parsing UUID string: ", err)
			return
		}
		newMsg := models.NewMessage(uuid, msg.ReceiverUuid, msg.Message)
		err = services.SaveMessage(newMsg)
		if err != nil {
			fmt.Println("Error when saving new msg in the database: ", err)
			return
		}

		fmt.Println("Message received:", msg.Message)
		err = conn.WriteJSON(newMsg)
		if err != nil {
			fmt.Println("Error when writing msg: ", err)
			break
		}
	}

}
