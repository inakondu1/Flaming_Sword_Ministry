package handlers

import (
	"fmt"
	"net/http"

	"Flaming_Sword_Ministry/database"
)

func ViewUsersHandler(w http.ResponseWriter, r *http.Request) {

	users, err := database.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "=== REGISTERED USERS ===")

	for _, u := range users {
		fmt.Fprintf(w, "ID: %d | Name: %s | Phone: %s | Gender: %s | Role: %s\n",
			u.ID, u.FullName, u.Phone, u.Gender, u.Role)
	}
}
