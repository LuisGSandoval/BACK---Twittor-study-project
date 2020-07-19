package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisGSandoval/twittor/db"
)

// FindProfile allows us to extract values from the profile
func FindProfile(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) == 0 {
		http.Error(w, "Profile not identified", http.StatusBadRequest)
		return
	}

	profile, err := db.FetchProfileData(ID)

	if err != nil {
		http.Error(w, "Couldn't  find the profile"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
