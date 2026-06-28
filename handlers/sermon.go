package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

// AddSermonHandler displays the form and saves a sermon.
func AddSermonHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/add_sermon.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		sermon := models.Sermon{
			Title:      r.FormValue("title"),
			BibleVerse: r.FormValue("bible_verse"),
			References: r.FormValue("references"),
			Content:    r.FormValue("content"),
			Category:   r.FormValue("category"),
			Date:       r.FormValue("date"),
			CreatedBy:  r.FormValue("created_by"),
		}

		err := database.CreateSermon(sermon)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/sermons", http.StatusSeeOther)
	}
}

// ViewSermonsHandler displays all sermons.
func ViewSermonsHandler(w http.ResponseWriter, r *http.Request) {

	sermons, err := database.GetAllSermons()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/sermons.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, sermons)
}
