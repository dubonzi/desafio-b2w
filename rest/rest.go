package rest

import (
	"desafio-b2w/service"
	"encoding/json"
	"net/http"
)

// SendJSON encodes data as json and writes it on w.
func SendJSON(w http.ResponseWriter, data interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		return service.ErrInternal
	}
	return nil
}

// SendError encodes err as json and writes it on w.
//	If err is of type service.HTTPError, its status code will be used for the response,
//	otherwise, http.StatusInternalServerError is used instead.
func SendError(w http.ResponseWriter, err error) error {
	hte, ok := err.(*service.HTTPError)
	if !ok {
		hte = service.ErrInternal
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(hte.Status)
	encoder := json.NewEncoder(w)
	return encoder.Encode(hte)
}
