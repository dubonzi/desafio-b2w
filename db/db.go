package db

import (
	"context"
	"desafio-b2w/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	mDB    *mongo.Database

	// DBName is the database name.
	DBName string

	// DBUrl is the database connection string.
	DBUrl string
)

// Open creates the initial connection to the MongoDB instance.
func Open() {
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(DBUrl))
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

// Close closes the database connection.
func Close() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Fatal("db.Close", "client.Disconnect", err)
	}
}

func getDB() *mongo.Database {
	if mDB != nil {
		return mDB
	}
	if client == nil {
		Open()
	}
	mDB = client.Database(DBName)

	return mDB
}
