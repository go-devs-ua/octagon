package http

import (
	"encoding/json"
	"log"
	"net/http"
)

//Usually, you need to make a rollback mechanism. If there is an error, roll back the changes
//from the database. But we will not do that now. Therefore, we will not change the status code
// after an error occurs
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	errRespond := map[string]interface{}{
		"error": InternalServerError,
	}

	if data == nil {
		log.Println("Data can not be empty") // we need to log the error when we select the logger
		WriteJSONResponse(w, statusCode, errRespond)
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("could not encode json: %v", err) //we need to log the error when we select the logge
		WriteJSONResponse(w, statusCode, errRespond)
	}
}
