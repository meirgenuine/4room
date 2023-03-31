package handlers

import (
	"4room/database"
	"4room/models"
	"4room/templates"
	"net/http"
	"strconv"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := struct {
			Title string
		}{
			Title: "Create Post",
		}
		templates.RenderTemplate(w, "create_post", &data)
	} else if r.Method == http.MethodPost {
		// Handle post creation form submission
		title := r.FormValue("title")
		content := r.FormValue("content")
		categoryIDStr := r.FormValue("category_id")
		userID := 1 // TODO: Get userID from the session

		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		post := models.Post{
			Title:      title,
			Content:    content,
			CategoryID: int64(categoryID),
			UserID:     int64(userID),
		}

		err = models.CreatePost(database.DB, &post)
		if err != nil {
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func ViewPost(w http.ResponseWriter, r *http.Request) {
	// Handle post viewing
	postIDStr := r.URL.Query().Get("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	post, err := models.GetPostByID(database.DB, int64(postID))
	if err != nil {
		http.Error(w, "Error fetching post", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title string
		Post  *models.Post
	}{
		Title: post.Title,
		Post:  post,
	}

	templates.RenderTemplate(w, "view_post", &data)
}
