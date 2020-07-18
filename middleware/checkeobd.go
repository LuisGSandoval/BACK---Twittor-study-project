package middleware

import (
	"net/http"

	"github.com/LuisGSandoval/twittor/db"
)

// ChequeoDB is a function that allows me to check the database status
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if db.ConnectionCheck() == 0 {
			http.Error(w, "Conexión perdida con la DB.", http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)

	}

}
