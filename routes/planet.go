package routes

import (
	"encoding/json"
	"net/http"
	"teste-b2w/model"
	"teste-b2w/rest"
	"teste-b2w/service"

	"github.com/go-chi/chi"
)

func planets() chi.Router {
	r := chi.NewRouter()

	r.Get("/", listAllPlanetsHandler)
	r.Get("/{id}", findPlanetByIDHandler)
	r.Post("/", newPlanetHandler)

	return r
}

func listAllPlanetsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	plService := service.PlanetService{}
	planets, err := plService.List(name)
	if err != nil {
		rest.SendError(w, err)
		return
	}

	rest.SendJSON(w, planets)
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
