package service

import (
	"net/http"
)

// HTTPError is an error message intended to be sent to the api clients.
type HTTPError struct {
	Status  int
	Code    string
	Message string
}

// Error implements the Error interface.
func (e *HTTPError) Error() string {
	return e.Message
}

var (
	// ErrInternal represents an internal server error.
	ErrInternal = &HTTPError{Status: http.StatusInternalServerError, Code: "internal_server_error", Message: "An unexpected internal server error occurred."}
	// ErrBadRequest represents a generic bad request error.
	ErrBadRequest = &HTTPError{Status: http.StatusBadRequest, Code: "bad_request", Message: "The information sent was in an invalid format."}
	// ErrNotFound represents a not found error.
	ErrNotFound = &HTTPError{Status: http.StatusNotFound, Code: "not_found", Message: "The requested resource was not found."}

	// ErrInvalidID represents an invalid ID error.
	ErrInvalidID = &HTTPError{Status: http.StatusBadRequest, Code: "invalid.id", Message: "The given ID is invalid."}

	// ErrEmptyName represents an empty name error.
	ErrEmptyName = &HTTPError{Status: http.StatusBadRequest, Code: "name.empty", Message: "The name cannot be empty."}
	// ErrEmptyClimate represents an empty climate error.
	ErrEmptyClimate = &HTTPError{Status: http.StatusBadRequest, Code: "climate.empty", Message: "The climate cannot be empty."}
	// ErrEmptyTerrain represents an empty terrain error.
	ErrEmptyTerrain = &HTTPError{Status: http.StatusBadRequest, Code: "terrain.empty", Message: "The terrain cannot be empty."}

	// ErrDuplicatedPlanet represents an error for duplicated planets.
	ErrDuplicatedPlanet = &HTTPError{Status: http.StatusBadRequest, Code: "planet.duplicated", Message: "A planet with the given name already exists."}
)
