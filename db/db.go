package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbURL = "mongodb://localhost:27017/"

var client *mongo.Client

var mDB *mongo.Database

func getDB() *mongo.Database {
	if mDB != nil {
		return mDB
	}

	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		log.Fatal("[FATAL] Error creating MongoDB client: ", err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		log.Fatal("[FATAL] Error connecting to the MongoDB instance: ", err)
	}

	// Check the connection
	client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("[FATAL] Error testing the connection:", err)
	}

	mDB = client.Database("starwars")

	return mDB
}
