package auth

import (
	"net/http"
)

func AuthRequired(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "bitslab-session")
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Redirect(w, r, "/login", 302)
			return
		}
		next.ServeHTTP(w, r)
	})
}
