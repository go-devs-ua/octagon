package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// here we will be taking care of resp and stuff

func DecodeBody(req *http.Request, obj any) error {
	if err := json.NewDecoder(req.Body).Decode(&obj); err != nil {
		return fmt.Errorf("error occurred while decoding request: %w", err)
	}
	defer req.Body.Close()
	return nil
}
