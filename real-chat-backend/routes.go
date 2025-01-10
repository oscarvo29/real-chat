package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *Config) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/", app.Index)
	mux.Get("/active-users", app.GetActiveUsers)
	mux.Post("/login", app.Login)

	return mux
}
