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
	router.HandleFunc("/registration", middleware.DBCheck(routers.Registration)).Methods("POST")

	handler := cors.AllowAll().Handler(router)

	log.Println("Server running on port: ", config.PORT())
	log.Fatal(http.ListenAndServe(":"+config.PORT(), handler))

}
