package main

import (
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

	log.Fatal(srv.ListenAndServe())
}
