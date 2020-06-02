package db

import (
	"context"
	"desafio-b2w/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mDB *mongo.Database

	// client is the client connected to the MongoDB instance.
	client *mongo.Client

	// DBName is the database name.
	DBName string

	// DBUri is the database connection string.
	DBUri string
)

// Open creates the initial connection to the MongoDB instance.
func Open() *mongo.Client {
	if client != nil {
		return client
	}
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(DBUri))
	if err != nil {
		logger.Fatal("db.Open", "mongo.NewClient", err)
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal("db.Open", "client.Connect", err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Fatal("db.Open", "client.Ping", err)
	}

	return client
}

// Close closes the database connection.
func Close() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Fatal("db.Close", "client.Disconnect", err)
	}
}

// GetDB returns the database handle.
func GetDB() *mongo.Database {
	if mDB != nil {
		return mDB
	}

	if client == nil {
		Open()
	}

	mDB = client.Database(DBName)

	return mDB
}
