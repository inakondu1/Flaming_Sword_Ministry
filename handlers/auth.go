package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"Flaming_Sword_Ministry/database"
	"Flaming_Sword_Ministry/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// Show registration page
	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	// Handle registration form
	if r.Method == http.MethodPost {
		fmt.Println("Registration form submitted")

		fullname := r.FormValue("fullname")
		phone := r.FormValue("phone")
		gender := r.FormValue("gender")
		password := r.FormValue("password")
		confirmPassword := r.FormValue("confirm_password")

		// Validation
		if fullname == "" || phone == "" || gender == "" || password == "" || confirmPassword == "" {
			http.Error(w, "Please fill in all fields.", http.StatusBadRequest)
			return
		}

		if password != confirmPassword {
			http.Error(w, "Passwords do not match.", http.StatusBadRequest)
			return
		}

		// Create User object
		user := models.User{
			FullName: fullname,
			Phone:    phone,
			Gender:   gender,
			Password: password, // We will hash this later
			Role:     "member",
		}

		// Save user
		err := database.CreateUser(user)
		if err != nil {
			http.Error(w, "Registration failed: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "🎉 Registration Successful!")
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		phone := r.FormValue("phone")
		password := r.FormValue("password")

		user, err := database.GetUserByPhone(phone)
		if err != nil || user.Password != password {
			http.Error(w, "Invalid phone or password", http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "Welcome %s 🎉 Login successful!", user.FullName)
	}
}
