package main

import (
	"log"
	"net/http"

	"4room/config"
	"4room/database"
	"4room/handlers"
	"4room/middleware"
)

func main() {
	config.LoadConfig()
	database.Initialize()

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))

	// Register handlers
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/logout", handlers.Logout)
	mux.HandleFunc("/create_post", middleware.AuthMiddleware(handlers.CreatePost))
	mux.HandleFunc("/view_post", handlers.ViewPost)
	mux.HandleFunc("/add_comment", middleware.AuthMiddleware(handlers.AddComment))
	mux.HandleFunc("/filter", middleware.AuthMiddleware(handlers.Filter))
	mux.HandleFunc("/list_categories", handlers.ListCategories)
	mux.HandleFunc("/create_category", middleware.AuthMiddleware(handlers.CreateCategory))

	mux.Handle("/static/", http.StripPrefix("/static", fs))

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	log.Printf("Server started at http://%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
