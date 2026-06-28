package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
)

// AdminDashboardHandler displays the admin dashboard.
func AdminDashboardHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/admin_dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ViewUsersHandler displays all registered users.
func ViewUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "===== REGISTERED MEMBERS =====")

	for _, u := range users {
		fmt.Fprintf(
			w,
			"ID: %d | Name: %s | Phone: %s | Gender: %s | Role: %s\n",
			u.ID,
			u.FullName,
			u.Phone,
			u.Gender,
			u.Role,
		)
	}
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
