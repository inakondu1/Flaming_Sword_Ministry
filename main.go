package main

import (
	"log"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/handlers"
	"Flaming_Sword_Ministry/middleware"
)

func main() {

	database.ConnectDB()

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)

	http.HandleFunc("/admin",
		middleware.AdminOnly(handlers.AdminHandler),
	)

	http.HandleFunc("/admin/announcement",
		middleware.AdminOnly(handlers.CreateAnnouncementHandler),
	)

	log.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
