package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oscarvo29/real-chat-backend/repositories"
	"github.com/oscarvo29/real-chat-backend/utils"
)

type Config struct{}

func main() {
	repositories.GetConnection(utils.DSN)
	fmt.Println(utils.DSN)

	app := Config{}

	srv := &http.Server{
		Handler: app.Routes(),
		Addr:    ":80",
	}

	fmt.Println("Server is running on port :80")
	log.Fatal(srv.ListenAndServe())
}
