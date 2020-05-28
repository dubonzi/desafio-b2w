package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Planet represents a planet from the Star Wars universe.
type Planet struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Name            string             `json:"name"`
	Climate         string             `json:"climate"`
	Terrain         string             `json:"terrain"`
	FilmAppearances int                `json:"filmAppearances"`
}
