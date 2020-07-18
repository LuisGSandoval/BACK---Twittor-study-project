package routers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/LuisGSandoval/twittor/db"
	"github.com/LuisGSandoval/twittor/models"
)

// Registration is the function that registers new users
func Registration(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	// sanitazing data
	t.Email = strings.TrimSpace(strings.ToLower(t.Email))

	// Error checking
	if err != nil {
		http.Error(w, "Error in the data: \n"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password must have more than 6 characters", http.StatusBadRequest)
		return
	}

	// checking before saving
	_, found, _ := db.UserExistanceCheck(t.Email)

	if found == true {
		http.Error(w, "This user already exists", http.StatusBadRequest)
		return
	}

	// Saving data
	_, status, err := db.UserResgistration(t)

	// Posible responses
	if err != nil {
		http.Error(w, "An error was encountered when saving the user. "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "User could not be saved", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
