package tests

import (
	"context"
	"desafio-b2w/db"
	"desafio-b2w/model"
	"desafio-b2w/routes"
	"desafio-b2w/service"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPlanetHandlers(t *testing.T) {
	setupDBVars()
	db.DBName = testDBName
	db.DBUri = testDBUri

	db.Open()

	t.Run("NewPlanet", newPlanet)
	t.Run("ListAllPlanets", listAllPlanets)
	t.Run("FindPlanetByID", findPlanetByID)
	t.Run("DeletePlanet", deletePlanet)

	db.Close()
}

func newPlanet(t *testing.T) {
	db.Client.Database(db.DBName).Collection("planets").Drop(context.Background())

	testPlanet := `{
		"name":    "Naboo",
		"terrain": "grassy hills, swamps, forests, mountains",
		"climate": "temperate"
	}`

	req, err := http.NewRequest("POST", "/api/planets", strings.NewReader(testPlanet))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.NewPlanetHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("wrong status returned by handler: expected %v but got %v", http.StatusCreated, status)
	}

	// From http://swapi.dev/api/planets/8/
	expectedFilms := 4

	var inserted model.Planet
	jsonDec := json.NewDecoder(recorder.Body)
	err = jsonDec.Decode(&inserted)
	if err != nil {
		t.Fatal("unable to parse response from handler, invalid body:", err)
	}
	if inserted.ID.IsZero() {
		t.Error("expected ID to be a non zero value")
	}
	if expectedFilms != inserted.FilmAppearances {
		t.Errorf("wrong amount of film appearances: expected %v but got %v", expectedFilms, inserted.FilmAppearances)
	}

}

func listAllPlanets(t *testing.T) {
	db.Client.Database(db.DBName).Collection("planets").Drop(context.Background())

	_, err := insertTestPlanetData(model.Planet{
		Name:    "Kamino",
		Climate: "temperate",
		Terrain: "ocean",
	})
	if err != nil {
		log.Fatal("error inserting test data: ", err)
	}
	_, err = insertTestPlanetData(model.Planet{
		Name:    "Alderaan",
		Climate: "temperate",
		Terrain: "grasslands, mountains",
	})
	if err != nil {
		log.Fatal("error inserting test data: ", err)
	}

	req, err := http.NewRequest("GET", "/api/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.ListPlanetsHandler)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status returned by handler: expected %v but got %v", http.StatusOK, status)
	}

	var planets []model.Planet
	jsonDec := json.NewDecoder(recorder.Body)
	err = jsonDec.Decode(&planets)
	if err != nil {
		t.Fatal("unable to parse response from handler, invalid body: ", err)
	}

	expectedPlanets := 2

	if len(planets) != expectedPlanets {
		t.Fatalf("wrong amount of planets returned from handler, expected %v but got %v", expectedPlanets, len(planets))
	}

	firstPlanetName := "Alderaan"

	if planets[0].Name != firstPlanetName {
		t.Errorf("wrong name of the first planet, expected %v but got %v", firstPlanetName, planets[0].Name)
	}

}

func findPlanetByID(t *testing.T) {
	db.Client.Database(db.DBName).Collection("planets").Drop(context.Background())

	insertedID, err := insertTestPlanetData(model.Planet{
		Name:    "Kamino",
		Climate: "temperate",
		Terrain: "ocean",
	})
	if err != nil {
		log.Fatal("error inserting test data: ", err)
	}

	req, err := http.NewRequest("GET", "/api/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.FindPlanetByIDHandler)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", insertedID.Hex())

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status returned by handler: expected %v but got %v", http.StatusOK, status)
	}

	var planet model.Planet
	jsonDec := json.NewDecoder(recorder.Body)
	err = jsonDec.Decode(&planet)
	if err != nil {
		t.Fatal("unable to parse response from handler, invalid body: ", err)
	}

	expectedName := "Kamino"

	if planet.Name != expectedName {
		t.Errorf("wrong planet returned from handler, expected name to be %v but got %v", expectedName, planet.Name)
	}

}

func deletePlanet(t *testing.T) {
	db.Client.Database(db.DBName).Collection("planets").Drop(context.Background())

	insertedID, err := insertTestPlanetData(model.Planet{
		Name:    "Kamino",
		Climate: "temperate",
		Terrain: "ocean",
	})
	if err != nil {
		log.Fatal("error inserting test data: ", err)
	}

	req, err := http.NewRequest("DELETE", "/api/planets", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.FindPlanetByIDHandler)

	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add("id", insertedID.Hex())

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx))

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("wrong status returned by handler: expected %v but got %v", http.StatusOK, status)
	}
}

// insertTestPlanetData inserts a planet into the test database.
func insertTestPlanetData(planet model.Planet) (primitive.ObjectID, error) {
	pService := service.PlanetService{}
	inserted, err := pService.Add(planet)
	return inserted.ID, err
}
