package db

import (
	"context"
	"teste-b2w/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const mongodbURL = "mongodb://localhost:27017/"

var client *mongo.Client

var mDB *mongo.Database

// Open creates the initial connection to the MongoDB instance.
func Open() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(mongodbURL))
	if err != nil {
		logger.Fatal("db.Open", "mongo.NewClient", err)
	}

	err = client.Connect(context.TODO())
	if err != nil {
		logger.Fatal("db.Open", "client.Connect", err)
	}

	// Check the connection
	client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal("db.Open", "client.Ping", err)
	}
}

func getDB() *mongo.Database {
	if mDB != nil {
		return mDB
	}
	if client == nil {
		Open()
	}
	mDB = client.Database("starwars")

	return mDB
}
