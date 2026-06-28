package handlers

import (
	"html/template"
	"net/http"
)

// HomeData holds the information shown on the homepage.
type HomeData struct {
	ChurchName string
	PastorName string
	Verse      string
	Prayer     string
}

// HomeHandler displays the homepage.
func HomeHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := HomeData{
		ChurchName: "Flaming Sword Ministry International",
		PastorName: "Pastor Pentecost",
		Verse:      "John 3:16",
		Prayer:     "Father, let Your fire consume every obstacle before me.",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/about.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
