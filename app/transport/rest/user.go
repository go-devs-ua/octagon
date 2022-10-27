package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-devs-ua/octagon/app/usecase"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUserResponse will wrap message
// that will be sent in JSON format.
type CreateUserResponse struct {
	ID string `json:"id"`
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
func (uh UserHandler) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var params = map[string]string{
			sort:   firstName + "," + lastName,
			offset: "",
			limit:  "",
		}

		for k, v := range req.URL.Query() {
			params[k] = strings.Join(v, "")
		}

		log.Printf("%+v", params)
		users, err := uh.usecase.Fetch(params[offset], params[limit], params[sort])
		if err != nil {
			uh.logger.Errorf("Failed fetching users from repository: %+v", err)
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not fetch users"}, uh.logger)

			return
		}

		res := struct {
			Results []entities.User `json:"results"`
		}{Results: users}

		WriteJSONResponse(w, http.StatusOK, res, uh.logger)
	})
}
