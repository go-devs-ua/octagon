package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GenerateJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data == nil {
		log.Println("Data can not be empty")
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Printf("could not encode json: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := "Internal Server Error"
		w.Write([]byte(errorMessage))
	}
}
