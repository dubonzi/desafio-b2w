package tests

import (
	"desafio-b2w/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type planetDBMock struct {
	existingPlanetName  string
	expectedPlanet      model.Planet
	existingPlanets     []model.Planet
	expectedAppearances int
}

// List lists planets.
func (p planetDBMock) List() ([]model.Planet, error) {
	return p.existingPlanets, nil
}

// Exists checks if a planet with the given name already exists
func (p planetDBMock) Exists(name string) (bool, error) {
	return p.existingPlanetName == name, nil
}

// Insert inserts a new planet into the collection.
func (p *planetDBMock) Insert(planet model.Planet) (primitive.ObjectID, error) {
	p.expectedPlanet = planet
	p.expectedPlanet.ID = primitive.NewObjectID()
	return p.expectedPlanet.ID, nil
}

// FindByID finds a planet by its id.
func (p planetDBMock) FindByID(id primitive.ObjectID) (model.Planet, error) {
	return p.expectedPlanet, nil
}

// FindByName finds a planet by its name (exact match).
func (p planetDBMock) FindByName(name string) (model.Planet, error) {
	return p.expectedPlanet, nil
}

// Delete deletes a planet.
func (p *planetDBMock) Delete(id primitive.ObjectID) error {
	foundIndex := -1
	for i, pl := range p.existingPlanets {
		if pl.ID.Hex() == id.Hex() {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return mongo.ErrNoDocuments
	}
	p.existingPlanets = p.existingPlanets[:len(p.existingPlanets)-1]
	return nil
}

// GetFilmAppearances returns the configured amount of appearances.
func (p planetDBMock) GetFilmAppearances(name string) (int, error) {
	return p.expectedAppearances, nil
}
