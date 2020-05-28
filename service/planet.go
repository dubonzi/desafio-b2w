package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"teste-b2w/db"
	"teste-b2w/logger"
	"teste-b2w/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PlanetService is the service layer structure that hold methods related to the Planet entity.
type PlanetService struct{}

// List lists planets with an optional name filter.
func (PlanetService) List(name string) ([]model.Planet, error) {
	plDB := db.NewPlanetDB()
	planets, err := plDB.List(name)
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
		logger.Error("PlanetService.FindByID", "plDB.FindByID", err, oID)
		return model.Planet{}, ErrInternal
	}

	return planet, nil
}

// Add adds a new planet.
//
// Must be a valid planet from the Star Wars universe.
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

	appearances, err := ps.getFilmAppearances(planet.Name)
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

func (PlanetService) getFilmAppearances(name string) (int, error) {
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
