package service

import (
	"log"
	"teste-b2w/db"
	"teste-b2w/model"
)

// PlanetService is the service layer structure that hold methods related to the Planet entity.
type PlanetService struct{}

// List lists planets with an optional name filter.
func (PlanetService) List() ([]model.Planet, error) {
	// TODO: Add filter
	plDB := db.NewPlanetDB()
	planets, err := plDB.List()
	if err != nil {
		// TODO: Improve logging
		log.Println("[ERROR] Error listing planets: ", err)
		return nil, ErrInternal
	}

	return planets, nil
}
