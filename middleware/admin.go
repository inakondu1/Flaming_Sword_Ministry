package middleware

import (
	"net/http"
)

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := Store.Get(r, "church-session")

		role, ok := session.Values["role"].(string)

		if !ok || role != "admin" {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}
