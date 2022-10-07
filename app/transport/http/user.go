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
			WriteJSONResponse(w, http.StatusBadRequest, Response{err.Error()})
			return
		}

		if err := uh.usecase.Signup(user); err != nil {
			// It is driving me nuts, please help. We need to decouple from pq package
			if err, ok := errors.Unwrap(err).(*pq.Error); ok && err.Code.Name() == "unique_violation" {
				WriteJSONResponse(w, http.StatusConflict, Response{MsgEmailConflict})
				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{MsgInternalSeverErr})
			return
		}
		// we could pass user here
		WriteJSONResponse(w, http.StatusCreated, Response{MsgUserCreated})
	})
}
