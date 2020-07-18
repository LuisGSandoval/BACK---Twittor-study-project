package db

import (
	"context"
	"time"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UserExistanceCheck checks if the user already exists to avoid duplicity
func UserExistanceCheck(email string) (models.User, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database(config.BDONE())

	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)

	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}

	return result, true, ID

}
