package db

import (
	"context"
	"time"

	"github.com/LuisGSandoval/twittor/config"
	"github.com/LuisGSandoval/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserResgistration this function saves new users in the Database
func UserResgistration(u models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database(config.BDONE())

	col := db.Collection("users")

	u.Password, _ = EncrypPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
