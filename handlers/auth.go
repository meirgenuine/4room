package handlers

import (
	"4room/database"
	"4room/models"
	"4room/templates"
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.RenderTemplate(w, "register", nil)
	case "POST":
		email := r.FormValue("email")
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Check if email is already taken
		var existingUser models.User
		err := database.DB.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&existingUser)
		if err == nil {
			http.Error(w, "Email is already taken", http.StatusConflict)
			return
		}

		// Hash the password
		passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}

		// Store the user in the database
		_, err = database.DB.Exec("INSERT INTO users (email, username, password_hash) VALUES (?, ?, ?)", email, username, passwordHash)
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Redirect to login page
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.RenderTemplate(w, "login", nil)
	case "POST":
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Retrieve the user from the database
		var user models.User
		err := database.DB.QueryRow("SELECT * FROM users WHERE email = ?", email).Scan(&user.ID, &user.Email, &user.Username, &user.PasswordHash)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Compare the provided password with the stored hash
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Generate session token
		sessionToken := make([]byte, 32)
		_, err = rand.Read(sessionToken)
		if err != nil {
			http.Error(w, "Error generating session token", http.StatusInternalServerError)
			return
		}

		// Set session cookie
		cookie := &http.Cookie{
			Name:     "session_token",
			Value:    base64.StdEncoding.EncodeToString(sessionToken),
			HttpOnly: true,
			Path:     "/",
			MaxAge:   86400, // 1 day
		}
		http.SetCookie(w, cookie)

		// Redirect to the main page
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Delete session cookie
	cookie := &http.Cookie{
		Name:   "session_token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	// Redirect to the main page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
