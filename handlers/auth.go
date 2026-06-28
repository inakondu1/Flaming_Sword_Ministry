package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/middleware"
	"Flaming_Sword_Ministry/models"
)

// REGISTER
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("templates/register.html")
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		r.ParseForm()

		user := models.User{
			FullName: r.FormValue("fullname"),
			Phone:    r.FormValue("phone"),
			Gender:   r.FormValue("gender"),
			Password: r.FormValue("password"),
			Role:     "member",
		}

		err := database.CreateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// LOGIN
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		r.ParseForm()

		phone := r.FormValue("phone")
		password := r.FormValue("password")

		user, err := database.GetUserByPhone(phone)
		if err != nil || user.ID == 0 {
			http.Error(w, "Invalid login", http.StatusUnauthorized)
			return
		}

		if user.Password != password {
			http.Error(w, "Invalid login", http.StatusUnauthorized)
			return
		}

		// SESSION START
		session, _ := middleware.Store.Get(r, "church-session")

		session.Values["user_id"] = user.ID
		session.Values["name"] = user.FullName
		session.Values["role"] = user.Role

		session.Save(r, w)

		fmt.Fprintf(w, "Welcome %s 🎉 Login successful!", user.FullName)
		return
	}
}
