package http

import (
	"encoding/json"
	"log"
	"net/http"
)

//Usually, you need to make a rollback mechanism. If there is an error, roll back the changes
//from the database. But we will not do that now. Therefore, we will not change the status code
// after an error occurs
func WriteJSONResponse(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	if data == nil {
		log.Println("Data is empty") // we need to log the error when we select the logger
		return
	}

	if err := json.NewEncoder(rw).Encode(data); err != nil {
		log.Printf("could not encode json: %v", err) //we need to log the error when we select the logge
	}
}