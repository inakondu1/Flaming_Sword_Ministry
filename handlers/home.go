package handlers

import (
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/middleware"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	name, ok := session.Values["name"].(string)

	data := map[string]interface{}{
		"LoggedIn": ok,
		"Name":     name,
	}

	tmpl, _ := template.ParseFiles("templates/home.html")
	tmpl.Execute(w, data)
}
