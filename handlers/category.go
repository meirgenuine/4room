package handlers

import (
	"4room/database"
	"4room/models"
	"4room/templates"
	"encoding/json"
	"net/http"
)

func ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories(database.DB)
	if err != nil {
		http.Error(w, "Error fetching categories: "+err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, "Error marshaling categories: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		templates.RenderTemplate(w, "create_category", nil)
		return
	} else if r.Method == "POST" {
		categoryName := r.FormValue("name")
		if categoryName == "" {
			http.Error(w, "Category name is required", http.StatusBadRequest)
			return
		}

		category := &models.Category{Name: categoryName}
		err := models.CreateCategory(database.DB, category)
		if err != nil {
			http.Error(w, "Error creating category: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
