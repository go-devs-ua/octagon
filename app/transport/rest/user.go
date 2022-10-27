package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-devs-ua/octagon/app/usecase"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUserResponse will wrap message
// that will be sent in JSON format.
type CreateUserResponse struct {
	ID string `json:"id"`
}

// User is entities.User version without sensitive data
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

// CreateUser will handle user creation.
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var user entities.User

		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)
			uh.logger.Errorf("Failed decoding JSON from request %+v: %+v", req, err)

			return
		}

		defer func() {
			if err := req.Body.Close(); err != nil {
				uh.logger.Warnf("Failed closing request %+v: %+v", req, err)
			}
		}()

		if err := user.Validate(); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)
			uh.logger.Errorf("Failed validating user: %+v", err)

			return
		}

		id, err := uh.usecase.Signup(user)
		if err != nil {
			uh.logger.Errorf("Failed creating user: %+v", err)

			if errors.Is(err, usecase.ErrDuplicateEmail) {
				WriteJSONResponse(w, http.StatusConflict, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, uh.logger)

			return
		}

		WriteJSONResponse(w, http.StatusCreated, CreateUserResponse{ID: id}, uh.logger)
		uh.logger.Debugw("User successfully created", "ID", id)
	})
}

// GetUsers retrieves all entities.User by given parameters.
//
// `offset` specified in request, if none was specified, default is defaultOffset
// `limit` specified in request, if none was specified, maxLimit is pagination limit default that is also request limit.
// `next` link is present if only `offset` and `limit` is less than total number of objects.
// defaultSort will be used if no params were specified.
func (uh UserHandler) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	})
}
