package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/services"
	"github.com/oscarvo29/real-chat-backend/utils"
)

func NewMessage(w http.ResponseWriter, r *http.Request) {
	senderUuid, ok := r.Context().Value(utils.UuidKey).(string)
	if !ok {
		http.Error(w, "Something went wrong with the server, when trying to parse the uuid.", http.StatusInternalServerError)
	}

	if senderUuid == "" {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
	}

	var newMsg models.Message
	err := utils.ParseJsonObject(r.Body, &newMsg)
	if err != nil {
		log.Fatal(err)
	}

	newMsg.SenderUuid, err = uuid.Parse(senderUuid)
	if err != nil {
		http.Error(w, "user id is not correct.", http.StatusUnauthorized)
	}

	err = services.SaveMessage(&newMsg)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}

	jsonData, err := json.Marshal(newMsg)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}
}

func ChatHistory(w http.ResponseWriter, r *http.Request) {
	senderUuidString, ok := r.Context().Value(utils.UuidKey).(string)
	if !ok {
		http.Error(w, "Something went wrong with the server, when trying to parse the uuid.", http.StatusInternalServerError)
	}

	if senderUuidString == "" {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
	}

	senderUuid, err := uuid.Parse(senderUuidString)
	if err != nil {
		http.Error(w, "user id is not correct.", http.StatusUnauthorized)
	}

	var chatHistoryRequest models.ChatHistory
	err = utils.ParseJsonObject(r.Body, &chatHistoryRequest)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}



	chatHistory, err := services.GetAllChatMessages(senderUuid, chatHistoryRequest.ReceiverUuid)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}
	jsonData, err := json.Marshal(chatHistory)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}
}
