package middleware

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("flaming-sword-secret-2026"))