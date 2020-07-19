package handlers

import (
	"log"
	"net/http"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/middleware"
	"github.com/LuisGSandoval/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores es el manejadr de rutas
func Manejadores() {
	router := mux.NewRouter()

	// All routers can be found here
	// Public
	router.HandleFunc("/api/registration", middleware.DBCheck(routers.Registration)).Methods("POST")
	router.HandleFunc("/api/login", middleware.DBCheck(routers.Login)).Methods("POST")

	// private
	router.HandleFunc("/api/profile", middleware.DBCheck(middleware.ValidateJWT(routers.FindProfile))).Methods("GET")

	// CORS config
	handler := cors.AllowAll().Handler(router)

	log.Println("Server running on port: ", config.Get("PORT"))
	log.Fatal(http.ListenAndServe(":"+config.Get("PORT"), handler))

}
