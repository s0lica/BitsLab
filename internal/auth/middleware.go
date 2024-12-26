package auth

import (
	"fmt"
	"net/http"

	"github.com/s0lica/BitsLab/internal/db"
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

func AdminRequired(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "bitslab-session")
		auth, ok := session.Values["authenticated"].(bool)
		if !auth || !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		db.InitDB()
		username := session.Values["username"]
		query := fmt.Sprintf("SELECT IsAdmin FROM Users WHERE Username='%s'", (username))
		var isadmin bool
		err := db.DB.QueryRow(query).Scan(&isadmin)
		if err != nil {
			panic(err)
		}
		db.CloseDB()
		if isadmin {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/", 302)
		}
	})

}
