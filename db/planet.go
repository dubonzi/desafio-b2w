package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// PlanetDB is a structure to acces 'planet' database records.
type PlanetDB struct {
	collection *mongo.Collection
}

// NewPlanetDB creates a new PlanetDB.
func NewPlanetDB() PlanetDB {
	p := PlanetDB{getDB().Collection("planets")}
	return p
}
