package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oscarvo29/real-chat-backend/models"
	"github.com/oscarvo29/real-chat-backend/services"
	"github.com/oscarvo29/real-chat-backend/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	err := utils.ParseJsonObject(r.Body, &usr)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte("user-not-found"))
	if err != nil {
		panic(err)
	}
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sign up have been hit !!! ")
	var newUsr models.User
	err := utils.ParseJsonObject(r.Body, &newUsr)
	if err != nil {
		log.Fatal(err)
	}

	err = services.SaveUser(&newUsr)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(fmt.Sprintf("User was created with ID: %v", newUsr.Uuid)))
	if err != nil {
		log.Fatal(err)
	}
}
