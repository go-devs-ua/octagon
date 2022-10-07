package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-devs-ua/octagon/lgr"
)

// Response will wrap message
// that will be sent in JSON format
type Response struct {
	Message string `json:"message"`
}

// TODO: Usually, you need to make a rollback mechanism.
//  If there is an error, roll back the changes
//  from the database. But we will not do that now.
//  Therefore, we will not change the status code
//  after an error occurs

// WriteJSONResponse writes JSON response
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data any, logger *lgr.Logger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		logger.Warnf("%s\n", "Data is empty")
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Errorf("could not encode json: %v", err)
	}
}
