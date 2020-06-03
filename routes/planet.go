package routes

import (
	"desafio-b2w/model"
	"desafio-b2w/rest"
	"desafio-b2w/service"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func planets() chi.Router {
	r := chi.NewRouter()

	r.Get("/", ListPlanetsHandler)
	r.Get("/{id}", FindPlanetByIDHandler)
	r.Post("/", NewPlanetHandler)
	r.Delete("/{id}", DeletePlanetHandler)

	return r
}

// ListPlanetsHandler handles requests for listing/searching planets.
func ListPlanetsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	plService := service.NewPlanetService(nil)

	var err error
	var searchResult interface{}

	if name != "" {
		searchResult, err = plService.FindByName(name)
	} else {
		searchResult, err = plService.List()
	}
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.SendJSON(w, searchResult, http.StatusOK)
}

// FindPlanetByIDHandler handles requests for finding planets by id.
func FindPlanetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	plService := service.NewPlanetService(nil)
	planet, err := plService.FindByID(id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.SendJSON(w, planet, http.StatusOK)
}

// NewPlanetHandler handles requests for inserting new planets.
func NewPlanetHandler(w http.ResponseWriter, r *http.Request) {
	var planet model.Planet
	jDec := json.NewDecoder(r.Body)
	err := jDec.Decode(&planet)
	if err != nil {
		rest.SendError(w, service.ErrBadRequest)
		return
	}
	plService := service.NewPlanetService(nil)
	planet, err = plService.Add(planet)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	rest.SendJSON(w, planet, http.StatusCreated)
}

// DeletePlanetHandler handles requests for deleting planets.
func DeletePlanetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	plService := service.NewPlanetService(nil)
	err := plService.Delete(id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}
