package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
)

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func CreateAnnouncementHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/create_announcement.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		r.ParseForm()

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
