package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/oscarvo29/real-chat-backend/models"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login have been hit")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var usr models.User
	err = json.Unmarshal(body, &usr)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("user-not-found"))
	if err != nil {
		panic(err)
	}
}

