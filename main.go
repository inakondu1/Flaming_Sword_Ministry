package main

import (
	"log"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/handlers"
)

func main() {

	// Connect to the database
	database.ConnectDB()

	// Serve static files
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// =========================
	// Public Routes
	// =========================
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	// Authentication
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)

	// Sermons
	http.HandleFunc("/sermons", handlers.ViewSermonsHandler)
	http.HandleFunc("/admin/add-sermon", handlers.AddSermonHandler)

	// Admin
	http.HandleFunc("/admin/users", handlers.ViewUsersHandler)
	http.HandleFunc("/admin/dashboard", handlers.AdminDashboardHandler)

	log.Println("🚀 Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}