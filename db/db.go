package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbURL = "mongodb://localhost:27017/"

var client *mongo.Client

// DB returns a connection to the database.
//	Creates a new one and returns it, if one does not already exist.
func DB() *mongo.Client {
	if client != nil {
		return client
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		log.Fatal("[FATAL] Error creating MongoDB client: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer client.Disconnect(ctx)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("[FATAL] Error connecting to the MongoDB instance: ", err)
	}

	return client
}
