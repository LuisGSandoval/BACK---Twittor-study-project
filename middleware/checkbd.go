package middleware

import (
	"net/http"

	"github.com/LuisGSandoval/twittor/db"
)

// DBCheck is a function that allows me to check the database status
func DBCheck(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if db.ConnectionCheck() == 0 {
			http.Error(w, "DB Connection lost.", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)

	}

}
