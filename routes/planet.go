package routes

import (
	"net/http"
	"teste-b2w/rest"
	"teste-b2w/service"

	"github.com/go-chi/chi"
)

func planets() chi.Router {
	r := chi.NewRouter()

	r.Get("/", listAllPlanetsHandler)

	return r
}

func listAllPlanetsHandler(w http.ResponseWriter, r *http.Request) {

	plService := service.PlanetService{}
	planets, err := plService.List()
	if err != nil {
		rest.SendError(w, err)
		return
	}

	rest.SendJSON(w, planets)
}
