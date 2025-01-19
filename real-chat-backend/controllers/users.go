package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/oscarvo29/real-chat-backend/services"
	"github.com/oscarvo29/real-chat-backend/utils"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	uuid, ok := r.Context().Value(utils.UuidKey).(string)
	if !ok {
		http.Error(w, "Something went wrong with the server, when trying to parse the uuid.", http.StatusInternalServerError)
	}

	if uuid == "" {
		http.Error(w, "User Unathorized. Try to log in again.", http.StatusUnauthorized)
	}

	users, err := services.GetAllUsers(uuid)
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
