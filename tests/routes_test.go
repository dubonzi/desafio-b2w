package tests

import (
	"desafio-b2w/db"
	"testing"
)

func TestMain(m *testing.M) {
	db.DBName = "starwars_test"
	db.DBUrl = "mongodb://localhost:27017/"
	db.Open()
	m.Run()
	db.Close()
}

func TestPlanetHandlers(t *testing.T) {

}
