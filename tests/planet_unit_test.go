// +build unit

package tests

import (
	"desafio-b2w/model"
	"desafio-b2w/service"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListPlanets(t *testing.T) {
	existingPlanets := []model.Planet{
		{
			Name:    "Naboo",
			Climate: "temperate",
			Terrain: "grassy hills, swamps, forests, mountains",
		}, {
			Name:    "Hoth",
			Climate: "frozen",
			Terrain: "tundra, ice caves, mountain ranges",
		},
	}

	mock := &planetDBMock{
		existingPlanets: existingPlanets,
	}

	pService := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier: mock,
	})

	foundPlanets, err := pService.List()

	if err != nil {
		t.Fatalf("expected err to be nil but got '%s'", err)
	}

	if len(existingPlanets) != len(foundPlanets) {
		// calling Fatal since an incorrect number of items might cause a panic in the next statement
		t.Fatalf("expected list of planets to have %d items, but it has %d", len(existingPlanets), len(foundPlanets))
	}

	if foundPlanets[0].Name != "Naboo" {
		t.Errorf("expected the first planet's name to be 'Naboo' but got %s", foundPlanets[0].Name)
	}
}

func TestPlanetExists(t *testing.T) {
	mock := &planetDBMock{
		existingPlanetName: "Naboo",
	}
	newPlanet := model.Planet{
		Name:    "Naboo",
		Climate: "temperate",
		Terrain: "grassy hills, swamps, forests, mountains",
	}

	ps := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier: mock,
	})
	_, err := ps.Add(newPlanet)
	if !errors.Is(err, service.ErrDuplicatedPlanet) {
		t.Errorf("expected '%s' error but got '%s' ", service.ErrDuplicatedPlanet, err)
	}
}

func TestPlanetValidation(t *testing.T) {
	mock := &planetDBMock{}
	newPlanet := model.Planet{}

	ps := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier: mock,
	})

	_, err := ps.Add(newPlanet)

	if !errors.Is(err, service.ErrEmptyName) {
		t.Errorf("expected '%s' error but got '%s' ", service.ErrEmptyName, err)
	}

	newPlanet.Name = "Tatooine"

	_, err = ps.Add(newPlanet)

	if !errors.Is(err, service.ErrEmptyClimate) {
		t.Errorf("expected '%s' error but got '%s' ", service.ErrEmptyClimate, err)
	}

	newPlanet.Climate = "arid"

	_, err = ps.Add(newPlanet)

	if !errors.Is(err, service.ErrEmptyTerrain) {
		t.Errorf("expected '%s' error but got '%s' ", service.ErrEmptyTerrain, err)
	}
}

func TestAddPlanet(t *testing.T) {
	mock := &planetDBMock{
		expectedAppearances: 7,
	}

	newPlanet := model.Planet{
		Name:    "Tatooine",
		Climate: "arid",
		Terrain: "desert",
	}

	pService := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier:            mock,
		GetFilmAppearances: mock.GetFilmAppearances,
	})

	addedPlanet, err := pService.Add(newPlanet)

	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}

	if addedPlanet.ID.IsZero() {
		t.Error("expected ID to have a value")
	}

	if addedPlanet.Name != newPlanet.Name {
		t.Errorf("expected Name to be %s but got %s", newPlanet.Name, addedPlanet.Name)
	}

	if addedPlanet.Climate != newPlanet.Climate {
		t.Errorf("expected Climate to be %s but got %s", newPlanet.Climate, addedPlanet.Climate)
	}

	if addedPlanet.Terrain != newPlanet.Terrain {
		t.Errorf("expected Terrain to be %s but got %s", newPlanet.Terrain, addedPlanet.Terrain)
	}
}

func TestFindPlanetByID(t *testing.T) {
	id := primitive.NewObjectID()
	expected := model.Planet{
		ID:      id,
		Name:    "Hoth",
		Climate: "frozen",
		Terrain: "tundra, ice caves, mountain ranges",
	}

	mock := &planetDBMock{
		expectedPlanet: expected,
	}

	pService := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier: mock,
	})

	foundPlanet, err := pService.FindByID(id.Hex())

	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}

	if foundPlanet.ID.Hex() != expected.ID.Hex() {
		t.Errorf("expected ID to be %s but got %s", expected.ID.Hex(), foundPlanet.ID.Hex())
	}

	if foundPlanet.Name != expected.Name {
		t.Errorf("expected Name to be %s but got %s", expected.Name, foundPlanet.Name)
	}
}

func TestFindPlanetByName(t *testing.T) {
	expected := model.Planet{
		Name:    "Hoth",
		Climate: "frozen",
		Terrain: "tundra, ice caves, mountain ranges",
	}

	mock := &planetDBMock{
		expectedPlanet: expected,
	}

	pService := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier: mock,
	})

	foundPlanet, err := pService.FindByName(expected.Name)

	if err != nil {
		t.Fatal("expected err to be nil but got: ", err)
	}

	if foundPlanet.Name != expected.Name {
		t.Errorf("expected name to be %s but got %s", expected.Name, foundPlanet.Name)
	}

}

func TestDeletePlanet(t *testing.T) {
	firstID, secondID := primitive.NewObjectID(), primitive.NewObjectID()

	// Keeping it simple, other properties are not needed.
	existingPlanets := []model.Planet{
		{
			ID:   firstID,
			Name: "Naboo",
		}, {
			ID:   secondID,
			Name: "Hoth",
		},
	}

	mock := &planetDBMock{
		existingPlanets: existingPlanets,
	}

	plService := service.NewPlanetService(&service.PlanetServiceOptions{
		Querier: mock,
	})

	err := plService.Delete(primitive.NewObjectID().Hex())

	if !errors.Is(err, service.ErrNotFound) {
		t.Errorf("expected err to be %s but got %s", service.ErrNotFound, err)
	}

	err = plService.Delete(secondID.Hex())

	if err != nil {
		t.Fatalf("expected err to be nil but got '%s'", err)
	}

	if len(existingPlanets) == len(mock.existingPlanets) {
		t.Errorf("expected list of planets to have less than %d items", len(existingPlanets))
	}
}
