package main

import (
	"log"

	"github.com/LuisGSandoval/twittor/db"
	"github.com/LuisGSandoval/twittor/handlers"
)

func main() {

	if db.ConnectionCheck() == 0 {
		log.Fatal("No connection to MongoDB")
		return
	}

	handlers.Manejadores()
}
