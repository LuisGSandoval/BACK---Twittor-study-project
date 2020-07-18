package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN is
var MongoCN = conectarBD()

var clientOptions = options.Client().ApplyURI("mongodb+srv://tuitor-app:1234@twittor-ver-1-baruh.mongodb.net/test?retryWrites=true&w=majority")

func conectarBD() *mongo.Client {

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

	log.Println("MongoDB exitosamente conectado")

	return client
}

// ConnectionCheck is
func ConnectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
