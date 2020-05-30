package db

import (
	"context"
	"desafio-b2w/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mDB *mongo.Database

	// Client is the client connected to the MongoDB instance.
	Client *mongo.Client

	// DBName is the database name.
	DBName string

	// DBUri is the database connection string.
	DBUri string
)

// Open creates the initial connection to the MongoDB instance.
func Open() *mongo.Client {
	if Client != nil {
		return Client
	}
	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI(DBUri))
	if err != nil {
		logger.Fatal("db.Open", "mongo.NewClient", err)
	}

	err = Client.Connect(context.TODO())
	if err != nil {
		logger.Fatal("db.Open", "Client.Connect", err)
	}

	// Check the connection
	Client.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal("db.Open", "Client.Ping", err)
	}

	return Client
}

// Close closes the database connection.
func Close() {
	err := Client.Disconnect(context.TODO())
	if err != nil {
		logger.Fatal("db.Close", "Client.Disconnect", err)
	}
}

func getDB() *mongo.Database {
	if mDB != nil {
		return mDB
	}

	if Client == nil {
		Open()
	}

	mDB = Client.Database(DBName)

	return mDB
}
