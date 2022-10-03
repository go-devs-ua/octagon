package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

// here we will be taking care of resp and stuff

func DecodeBody(r *http.Request, u *entities.User) error {
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		return err
	}
	return nil
}
