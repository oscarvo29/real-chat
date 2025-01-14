package middleware

import (
	"fmt"
	"net/http"

	"github.com/oscarvo29/real-chat-backend/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		uuid, err := utils.VerifyToken(authHeader)
		if err != nil {
			http.Error(w, "Client is not Authorized", http.StatusUnauthorized)
		}

		if uuid != "" {
			next.ServeHTTP(w, r)
		}
	})
}
