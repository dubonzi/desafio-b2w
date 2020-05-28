package db

import (
	"context"
	"teste-b2w/model"

	"go.mongodb.org/mongo-driver/bson"
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

// List lists planets.
func (p PlanetDB) List() ([]model.Planet, error) {
	cursor, err := p.collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	planets := make([]model.Planet, 0, 5)
	for cursor.Next(context.Background()) {
		var planet model.Planet
		err = cursor.Decode(&planet)
		if err != nil {
			return nil, err
		}
		planets = append(planets, planet)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.Background())
	return planets, nil
}
