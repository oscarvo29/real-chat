package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

func main() {
	app := Config{}

	srv := &http.Server{
		Handler: app.Routes(),
		Addr:    ":80",
	}

	fmt.Println("Server is running on port :80")
	log.Fatal(srv.ListenAndServe())
}
