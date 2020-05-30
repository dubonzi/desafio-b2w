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

	r.Get("/", listPlanetsHandler)
	r.Get("/{id}", findPlanetByIDHandler)
	r.Post("/", newPlanetHandler)
	r.Delete("/{id}", deletePlanetHandler)

	return r
}

func listPlanetsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	plService := service.PlanetService{}

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
	rest.SendJSON(w, searchResult)
}

func findPlanetByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	plService := service.PlanetService{}
	planet, err := plService.FindByID(id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
	rest.SendJSON(w, planet)
}

func newPlanetHandler(w http.ResponseWriter, r *http.Request) {
	var planet model.Planet
	jDec := json.NewDecoder(r.Body)
	err := jDec.Decode(&planet)
	if err != nil {
		rest.SendError(w, service.ErrBadRequest)
		return
	}
	plService := service.PlanetService{}
	planet, err = plService.Add(planet)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	rest.SendJSON(w, planet)
}

func deletePlanetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	plService := service.PlanetService{}
	err := plService.Delete(id)
	if err != nil {
		rest.SendError(w, err)
		return
	}
}
