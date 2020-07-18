package routers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LuisGSandoval/twittor/db"
	"github.com/LuisGSandoval/twittor/jwt"
	"github.com/LuisGSandoval/twittor/models"
)

// Login is the handler that recieves the login requests
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	// sanitazing data
	t.Email = strings.TrimSpace(strings.ToLower(t.Email))

	if err != nil {
		http.Error(w, "Error, something when wrong "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "email required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "password must have more than 6 characters", http.StatusBadRequest)
		return
	}

	document, succes := db.ValidatePasswords(t.Email, t.Password)

	if !succes {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Error encountered when generating new token"+err.Error(), http.StatusInternalServerError)
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Just in case we want to use cookies

	// expirationTime := time.Now().Add(24 * time.Hour)
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   jwtKey,
	// 	Expires: expirationTime,
	// })

}
