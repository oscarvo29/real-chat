package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oscarvo29/real-chat-backend/services"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		panic(err)
	}
}
