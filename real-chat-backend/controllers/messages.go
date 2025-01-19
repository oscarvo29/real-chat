package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/services"
	"github.com/oscarvo29/real-chat-backend/utils"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {
	senderUuidString, ok := r.Context().Value(utils.UuidKey).(string)
	if !ok {
		http.Error(w, "Something went wrong with the server, when trying to parse the uuid.", http.StatusInternalServerError)
	}

	if senderUuidString == "" {
		http.Error(w, "user not authorized", http.StatusUnauthorized)
	}

	var chat models.Chat
	err := utils.ParseJsonObject(r.Body, &chat)
	if err != nil {
		http.Error(w, "could not convert the ids.", http.StatusUnauthorized)
		return
	}
	chat.Participants = append(chat.Participants, senderUuidString)
	err = services.CreateChat(&chat)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "could not create the chat", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(chat)
	if err != nil {
		http.Error(w, "Problem converting chat into json data.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}

}

func GetAllChatsForUser(w http.ResponseWriter, r *http.Request) {
	senderUuid, err := utils.TransFormJWT(r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "user id is not correct.", http.StatusUnauthorized)
		return
	}

	chats, err := services.GetChatRooms(senderUuid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong when trying to load users chats", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(&chats)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}

}

func ChatHistory(w http.ResponseWriter, r *http.Request) {
	chatId := chi.URLParam(r, "chatId")
	senderUuid, err := utils.TransFormJWT(r)
	if err != nil {
		fmt.Println("err: ", err)
		http.Error(w, "There was a problem. Try to logout and in again.", http.StatusInternalServerError)
	}

	messages, err := services.GetChatHistory(chatId, senderUuid)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong when trying to load users chats", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(&messages)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}
}
