package model

// Planet represents a planet from the Star Wars universe.
type Planet struct {
	Name            string `json:"name"`
	Climate         string `json:"climate"`
	Terrain         string `json:"terrain"`
	FilmAppearences uint   `json:"filmAppearences"`
}
