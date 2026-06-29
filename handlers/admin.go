package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
)

type AdminData struct {
	Name               string
	TotalUsers         int
	TotalSermons       int
	TotalAnnouncements int
	Users              interface{}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)

	totalUsers, _ := database.CountUsers()
	totalSermons, _ := database.CountSermons()
	totalAnnouncements, _ := database.CountAnnouncements()

	users, _ := database.GetAllUsers()

	data := AdminData{
		Name:               name,
		TotalUsers:         totalUsers,
		TotalSermons:       totalSermons,
		TotalAnnouncements: totalAnnouncements,
		Users:              users,
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ================= VIEW USERS =================

func ViewUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/users.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
