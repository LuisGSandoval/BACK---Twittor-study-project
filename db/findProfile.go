package db

import (
	"context"
	"time"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FetchProfileData is a function that brings information about a specific profile
func FetchProfileData(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database(config.Get("DBONE"))

	col := db.Collection(config.Get("USRCOL"))

	var profile models.User

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)

	if err != nil {
		return profile, err
	}

	profile.Password = ""

	return profile, nil

}
