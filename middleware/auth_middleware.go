package middleware

import (
	"4room/models"
	"errors"
	"net/http"
)

func GetUserFromSession(r *http.Request) (*models.User, error) {
	// TODO: Implement session management and fetch the user from the session
	return nil, errors.New("Not implemented")
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := GetUserFromSession(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx := r.Context()
		ctx = models.NewUserContext(ctx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
