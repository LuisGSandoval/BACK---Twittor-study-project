package middleware

import (
	"net/http"

	"github.com/LuisGSandoval/twittor/routers"
)

// ValidateJWT validates the JWT and the user already has it
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		_, _, _, err := routers.ProcessToken(token)

		if err != nil {
			http.Error(w, "Token error: "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
