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

	passwordMatched, err := services.ValidateUser(&usr)
	if err != nil {
		panic(err)
	}

	if passwordMatched {
		jwt, err := utils.GenerateToken(usr.Name, usr.Uuid)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write([]byte(jwt))
		if err != nil {
			panic(err)
		}
		return
	}
	_, err = w.Write([]byte("Password wasn't correct."))
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
	jwt, err := utils.GenerateToken(newUsr.Name, newUsr.Uuid)
	if err != nil {
		panic(err)
	}
	_, err = w.Write([]byte(jwt))
	if err != nil {
		log.Fatal(err)
	}
}
