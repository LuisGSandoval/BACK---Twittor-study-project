package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LuisGSandoval/twittor/db"
	"github.com/LuisGSandoval/twittor/models"
)

// Registration es la ruta que regstra un usuario en la base de datos
func Registration(w http.ResponseWriter, r *http.Request) {

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos resibidos: \n"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email vacio", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Contraseña con menos de 6 caracteres", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := db.CheckIfUserExists(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe este usuario", http.StatusBadRequest)
		return
	}

	_, status, err := db.UserResgistration(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar el registro de usuario"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se pudo registrar el usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
