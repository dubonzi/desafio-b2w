package tests

import (
	"os"
)

var (
	testDBName string
	testDBUri  string
)

func setupDBVars() {
	testDBUri = os.Getenv("MONGO_TEST_DATABASE_URI")
	testDBName = os.Getenv("MONGO_TEST_DATABASE_NAME")
	if testDBName == "" {
		testDBName = "starwars_testdb"
	}
	if testDBUri == "" {
		testDBUri = "mongodb://localhost:27017/"
	}
}
