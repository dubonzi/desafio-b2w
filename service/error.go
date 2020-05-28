package service

import "net/http"

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
)
