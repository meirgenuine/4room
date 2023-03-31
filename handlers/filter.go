package handlers

import (
	"4room/database"
	"4room/models"
	"html/template"
	"net/http"
	"strconv"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		categoryIDStr := r.FormValue("category_id")
		categoryID, err := strconv.ParseInt(categoryIDStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid category ID", http.StatusBadRequest)
			return
		}

		posts, err := models.GetPostsByCategory(database.DB, categoryID)
		if err != nil {
			http.Error(w, "Error fetching posts", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/base.html", "templates/filter.html")
		if err != nil {
			http.Error(w, "Error loading templates", http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"Title": "Filtered Posts",
			"Posts": posts,
		}

		err = tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
