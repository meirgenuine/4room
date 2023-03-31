package handlers

import (
	"4room/database"
	"4room/models"
	"4room/templates"
	"net/http"
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
	}{
		Title: "4room - Home",
		Posts: posts,
	}

	templates.RenderTemplate(w, "index", &data)
}
