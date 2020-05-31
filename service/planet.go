package service

import (
	"desafio-b2w/db"
	"desafio-b2w/logger"
	"desafio-b2w/model"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// PlanetService is the service layer structure that hold methods related to the Planet entity.
type PlanetService struct{}

// List lists planets.
func (PlanetService) List() ([]model.Planet, error) {
	plDB := db.NewPlanetDB()
	planets, err := plDB.List()
	if err != nil {
		logger.Error("PlanetService.List", "plDB.List", err)
		return nil, ErrInternal
	}

	return planets, nil
}

// FindByID returns a planet with the given id.
//
// Returns ErrNotFound if the planet doesn't exist.
func (PlanetService) FindByID(id string) (model.Planet, error) {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Planet{}, ErrInvalidID
	}

	plDB := db.NewPlanetDB()
	planet, err := plDB.FindByID(oID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Planet{}, ErrNotFound
		}
		logger.Error("PlanetService.FindByID", "plDB.FindByID", err, oID)
		return model.Planet{}, ErrInternal
	}

	return planet, nil
}

// FindByName returns a planet with the given name.
//
// Returns ErrNotFound if the planet doesn't exist.
func (PlanetService) FindByName(name string) (model.Planet, error) {
	plDB := db.NewPlanetDB()
	planet, err := plDB.FindByName(name)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Planet{}, ErrNotFound
		}
		logger.Error("PlanetService.FindByName", "plDB.FindByName", err, name)
		return model.Planet{}, ErrInternal
	}
	return planet, nil
}

// Add adds a new planet.
func (ps PlanetService) Add(planet model.Planet) (model.Planet, error) {
	if planet.Name == "" {
		return model.Planet{}, ErrEmptyName
	}
	if planet.Climate == "" {
		return model.Planet{}, ErrEmptyClimate
	}
	if planet.Terrain == "" {
		return model.Planet{}, ErrEmptyTerrain
	}
	plDB := db.NewPlanetDB()

	exists, err := plDB.Exists(planet.Name)
	if err != nil {
		logger.Error("PlanetService.Add", "plDB.Exists", err, planet.Name)
		return model.Planet{}, ErrInternal
	}
	if exists {
		return model.Planet{}, ErrDuplicatedPlanet
	}

	appearances, err := ps.GetFilmAppearances(planet.Name)
	if err != nil {
		logger.Error("PlanetService.Add", "ps.getFilmAppearances", err, planet.Name)
		return model.Planet{}, ErrInternal
	}
	planet.FilmAppearances = appearances
	planet.ID = primitive.NewObjectID()
	id, err := plDB.Insert(planet)
	if err != nil {
		logger.Error("PlanetService.Add", "plDB.Insert", err, planet)
		return model.Planet{}, ErrInternal
	}

	planet, err = plDB.FindByID(id)
	if err != nil {
		logger.Error("PlanetService.Add", "plDB.FindByID", err, id)
		return model.Planet{}, ErrInternal
	}

	return planet, nil
}

// Delete removes a planet from the database.
func (PlanetService) Delete(id string) error {
	oID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidID
	}
	plDB := db.NewPlanetDB()
	err = plDB.Delete(oID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrNotFound
		}
		logger.Error("PlanetService.Delete", "plDB.Delete", err, oID)
		return ErrInternal
	}

	return nil
}

// GetFilmAppearances returns the number of films a planet with `name` appeared on.
// If multiple planets are found, the first planet from the list will be chosen.
// Returns 0 if no planets are found.
func (PlanetService) GetFilmAppearances(name string) (int, error) {
	client := http.Client{}
	resp, err := client.Get(fmt.Sprintf("https://swapi.dev/api/planets/?search=%s", name))
	if err != nil {
		return 0, err
	}

	// Using a anonymous struct since we only need information about Films from the response.
	var search struct {
		Results []struct {
			Films []string `json:"films"`
		} `json:"results"`
	}

	jsDec := json.NewDecoder(resp.Body)
	err = jsDec.Decode(&search)
	if err != nil {
		return 0, err
	}

	if len(search.Results) > 0 {
		// Using the first result because SWAPI's search uses "case-insensitive partial matches on the set of search fields".
		return len(search.Results[0].Films), nil
	}
	return 0, nil
}
