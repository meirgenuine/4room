package handlers

import (
	"net/http"

	"4room/database"
	"4room/models"
	"4room/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Handle the main page
	posts, err := models.GetAllPosts(database.DB)
	if err != nil {
		http.Error(w, "Error fetching posts", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Posts []models.Post
		User  *models.User
	}{
		Title: "4room - Home",
		Posts: posts,
		User:  models.UserFromContext(r.Context()),
	}

	templates.RenderTemplate(w, "index", &data)
}
