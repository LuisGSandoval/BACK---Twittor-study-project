package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/LuisGSandoval/twittor/middleware"
	"github.com/LuisGSandoval/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores es el manejadr de rutas
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.ChequeoDB(routers.Registration)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
