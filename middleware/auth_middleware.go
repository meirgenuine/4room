package middleware

import (
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if user is authenticated, if not redirect to login page
		// If authenticated, call the next handler
	}
}
