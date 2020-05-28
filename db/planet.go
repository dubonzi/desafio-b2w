package db

import (
	"context"
	"errors"
	"teste-b2w/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// List lists planets with an optional name filter.
func (p PlanetDB) List(name string) ([]model.Planet, error) {
	filter := bson.M{"name": bson.M{"$regex": ".*" + name + ".*"}}
	cursor, err := p.collection.Find(context.Background(), filter)
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

// Exists checks if a planet with the given name already exists
func (p PlanetDB) Exists(name string) (bool, error) {
	result := p.collection.FindOne(context.Background(), bson.M{"name": name})
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, result.Err()
	}
	return true, nil
}

// Insert inserts a new planet into the collection.
func (p PlanetDB) Insert(planet model.Planet) (primitive.ObjectID, error) {
	result, err := p.collection.InsertOne(context.Background(), planet)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

// FindByID finds a planet by its id.
func (p PlanetDB) FindByID(id primitive.ObjectID) (model.Planet, error) {
	result := p.collection.FindOne(context.Background(), bson.M{"_id": id})
	if result.Err() != nil {
		return model.Planet{}, result.Err()
	}
	var planet model.Planet
	return planet, result.Decode(&planet)
}

// Delete deletes a planet.
func (p PlanetDB) Delete(id primitive.ObjectID) error {
	return p.collection.FindOneAndDelete(context.Background(), bson.M{"_id": id}).Err()
}
