package conf

import "os"

var (
	mongoDBURI              = "mongodb://localhost:27017/"
	mongoDBDatabaseName     = "starwars"
	mongoDBTestDatabaseName = "starwars_testdb"
)

// Load config variables from the environment.
func Load() {
	if os.Getenv("MONGODB_URI") != "" {
		mongoDBURI = os.Getenv("MONGODB_URI")
	}
	if os.Getenv("MONGODB_DATABASE_NAME") != "" {
		mongoDBDatabaseName = os.Getenv("MONGODB_DATABASE_NAME")
	}
	if os.Getenv("MONGODB_TEST_DATABASE_NAME") != "" {
		mongoDBTestDatabaseName = os.Getenv("MONGODB_TEST_DATABASE_NAME")
	}

}

// MongoDBURI URI used to connect to the MongoDB instance.
//	Default: "mongodb://localhost:27017/"
func MongoDBURI() string {
	return mongoDBURI
}

// MongoDBDatabaseName name of the database.
//	Default: "starwars"
func MongoDBDatabaseName() string {
	return mongoDBDatabaseName
}

// MongoDBTestDatabaseName name of the test database.
//	Default: "starwars_testdb"
func MongoDBTestDatabaseName() string {
	return mongoDBTestDatabaseName
}
