package db

import (
	"context"
	"time"

	"github.com/LuisGSandoval/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// CheckIfUserExists checks if the user already exists to avoid duplicity
func CheckIfUserExists(email string) (models.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twitter")

	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.User

	err := col.FindOne(ctx, condition).Decode(&resultado)

	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
