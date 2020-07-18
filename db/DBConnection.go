package db

import (
	"context"
	"log"

	"github.com/LuisGSandoval/twittor/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is the connection to mongoDB
var MongoCN = connectBD()

var clientOptions = options.Client().ApplyURI(config.MONGODBCXSTRING())

func connectBD() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)

		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)

		return client
	}

	log.Println("MongoDB connected succesfully")

	return client
}

// ConnectionCheck is a test done to the database to make sure it's working as expected
func ConnectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
