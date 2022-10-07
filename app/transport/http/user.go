package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/lib/pq"
)

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		user := entities.User{}

		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{MsgBadRequest})
			return
		}

		defer func() {
			if err := req.Body.Close(); err != nil {
				log.Println(err)
			}
		}()

		if err := user.Validate(); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{"Validation error: " + err.Error()})
			return
		}

		if err := uh.usecase.Signup(user); err != nil {
			// TODO: Handle errors gracefully
			if err, ok := errors.Unwrap(err).(*pq.Error); ok && err.Code.Name() == "unique_violation" {
				WriteJSONResponse(w, http.StatusConflict, Response{MsgEmailConflict})
				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{MsgInternalSeverErr})
			return
		}

		WriteJSONResponse(w, http.StatusCreated, Response{MsgUserCreated})
	})
}
