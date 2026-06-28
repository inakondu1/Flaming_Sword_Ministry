package handlers

import (
	"net/http"

	"Flaming_Sword_Ministry/middleware"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := middleware.Store.Get(r, "church-session")

	delete(session.Values, "user_id")
	delete(session.Values, "name")
	delete(session.Values, "role")

	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
