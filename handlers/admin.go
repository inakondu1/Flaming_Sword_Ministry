package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
)

// AdminHandler displays the admin dashboard.
func AdminHandler(w http.ResponseWriter, r *http.Request) {

	// Get current session
	session, _ := middleware.Store.Get(r, "church-session")

	name, _ := session.Values["name"].(string)

	totalUsers, _ := database.CountUsers()
	totalSermons, _ := database.CountSermons()
	totalAnnouncements, _ := database.CountAnnouncements()

	data := struct {
		Name               string
		TotalUsers         int
		TotalSermons       int
		TotalAnnouncements int
	}{
		Name:               name,
		TotalUsers:         totalUsers,
		TotalSermons:       totalSermons,
		TotalAnnouncements: totalAnnouncements,
	}

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// CreateAnnouncementHandler displays the announcement form and saves announcements.
func CreateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {

	// Display the form
	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/announcement.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	// Save the announcement
	if r.Method == http.MethodPost {

		title := r.FormValue("title")
		message := r.FormValue("message")

		err := database.CreateAnnouncement(title, message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
