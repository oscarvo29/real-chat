package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/oscarvo29/real-chat-backend/controllers"
	"github.com/oscarvo29/real-chat-backend/middleware"
)

type contextKey string

const uuidKey contextKey = "uuid"

func (app *Config) Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Route("/auth", func(r chi.Router) {
		r.Post("/login", controllers.LoginHandler)
		r.Post("/signup", controllers.SignUpHandler)
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Route("/users", func(r chi.Router) {
			r.Get("/all-users", controllers.GetAllUsers)
		})
	})

	mux.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Route("/messages", func(r chi.Router) {
			r.Post("/send-message", controllers.NewMessage)
			r.Post("/get-chat-history", controllers.ChatHistory)

		})
	})

	mux.HandleFunc("/chat-ws", controllers.HandleConnections)

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("index page"))
	})

	return mux
}
