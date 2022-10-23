package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/gorilla/mux"
	"github.com/lib/pq"
)

// Response will wrap message
// that will be sent in JSON format
type CreateUserResponse struct {
	ID string `json:"id"`
}

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var user entities.User

		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: BadRequestMsg, Details: err.Error()}, uh.logger)
			uh.logger.Errorf("Failed decoding JSON from request %+v: %+v", req, err)
			return
		}

		defer func() {
			if err := req.Body.Close(); err != nil {
				uh.logger.Warnf("Failed closing request %+v: %+v", req, err)
			}
		}()

		if err := user.Validate(); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: BadRequestMsg, Details: err.Error()}, uh.logger)
			uh.logger.Errorf("Failed validating user: %+v", err)
			return
		}

		id, err := uh.usecase.Signup(user)
		if err != nil {
			uh.logger.Errorf("Failed creating user: %+v", err)

			if err, ok := errors.Unwrap(err).(*pq.Error); ok && err.Code.Name() == "unique_violation" {
				WriteJSONResponse(w, http.StatusConflict, Response{Message: BadRequestMsg, Details: err.Error()}, uh.logger)
				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: ServerErrMsg}, uh.logger)
			return
		}

		WriteJSONResponse(w, http.StatusCreated, CreateUserResponse{ID: id}, uh.logger)
		uh.logger.Debugw("User successfully created", "ID", id)
	})
}

func (uh UserHandler) GetUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]

		user, err := uh.usecase.GetUser(id)
		if err != nil {
			uh.logger.Errorf("Failed creating user: %+v", err)
			if err, ok := errors.Unwrap(err).(*pq.Error); ok && err.Code.Name() == "unique_violation" {
				WriteJSONResponse(w, http.StatusConflict, Response{MsgEmailConflict}, uh.logger)
				return
			}

			WriteJSONResponse(w, http.StatusNotFound, Response{MsgNotFound}, uh.logger)
			return
		}

		WriteJSONResponse(w, http.StatusCreated, user, uh.logger)
		uh.logger.Debugw("User received", "ID", id)
	})
}
