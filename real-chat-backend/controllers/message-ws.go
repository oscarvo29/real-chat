package controllers

import (
	"encoding/json"
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

type socketDataObj struct {
	Event string      `json:"event"`
	Jwt   string      `json:"jwt"`
	Data  interface{} `json:"data"`
}

var clients = make(map[string]*websocket.Conn)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Web socket have been hit?")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error when upgrading to websocket: ", err)
		return
	}

	for {
		var socketConn socketDataObj
		err := conn.ReadJSON(&socketConn)
		if err != nil {
			fmt.Println("Error when reading message: ", err)
			return
		}

		switch socketConn.Event {
		case "connection_open":

			uuid, err := utils.VerifyToken(socketConn.Jwt)
			if err != nil {
				fmt.Println("Error when converting JWT: ", err)
				conn.Close()
				return
			}
			clients[uuid] = conn
		case "message":
			var shortMsg models.ShortMessage
			jsonData, err := json.Marshal(socketConn.Data)
			if err != nil {
				fmt.Println("Error when trying to valite the data: ", err)
				conn.Close()
				return
			}

			err = json.Unmarshal(jsonData, &shortMsg)
			if err != nil {
				fmt.Println("Error when trying to valite the data after json convertion: ", err)
				conn.Close()
				return
			}

			uiidStr, err := utils.VerifyToken(shortMsg.SenderJWT)
			if err != nil {
				fmt.Println("error validating jwt token: ", err)
				conn.Close()
				return
			}

			senderId, err := uuid.Parse(uiidStr)
			if err != nil {
				fmt.Println("error converting uuid string: ", err)
				conn.Close()
				return
			}

			newMsg := models.NewMessage(senderId, shortMsg.ChatId, shortMsg.Message)
			err = services.SaveMessage(newMsg)
			if err != nil {
				fmt.Println("Error saving message: ", err)
				conn.Close()
				return
			}

			receipiants, err := services.GetReceipiants(newMsg.ChatId)
			if err != nil {
				fmt.Println("Error when trying to get receipiants for the message: ", err)
				conn.Close()
				return
			}
			for _, receipiant := range receipiants {
				key := *receipiant

				clientConn, ok := clients[key]
				if ok {
					if newMsg.SenderUuid.String() == key {
						newMsg.IsSender = true
					}

					err = clientConn.WriteJSON(newMsg)
					if err != nil {
						fmt.Println("Error when writing msg to receiver: ", err)
					}
					newMsg.IsSender = false
				}
			}
		case "conn_close":
			uuid, err := utils.VerifyToken(socketConn.Jwt)
			if err != nil {
				fmt.Println("Error when converting JWT: ", err)
				conn.Close()
				return
			}
			conn.Close()
			delete(clients, uuid)
		}

		// err = conn.WriteJSON(newMsg)
		// if err != nil {
		// 	fmt.Println("Error when writing msg: ", err)
		// 	break
		// }
	}

}

// func HandleConnections(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Web socket have been hit?")
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error when upgrading to websocket: ", err)
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		var msg models.ShortMessage
// 		err := conn.ReadJSON(&msg)
// 		if err != nil {
// 			fmt.Println("Error when reading message: ", err)
// 			return
// 		}

// 		uuidString, err := utils.VerifyToken(msg.SenderJWT)
// 		if err != nil {
// 			fmt.Println("Error when reading JWT: ", err)
// 			return
// 		}
// 		uuid, err := uuid.Parse(uuidString)
// 		if err != nil {
// 			fmt.Println("Error when parsing UUID string: ", err)
// 			return
// 		}
// 		newMsg := models.NewMessage(uuid, msg.ReceiverUuid, msg.Message)
// 		err = services.SaveMessage(newMsg)
// 		if err != nil {
// 			fmt.Println("Error when saving new msg in the database: ", err)
// 			return
// 		}

// 		fmt.Println("Message received:", msg.Message)
// 		err = conn.WriteJSON(newMsg)
// 		if err != nil {
// 			fmt.Println("Error when writing msg: ", err)
// 			break
// 		}
// 	}

// }
