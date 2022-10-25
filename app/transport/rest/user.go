package rest

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-devs-ua/octagon/app/usecase"
	"github.com/gorilla/mux"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUserResponse will wrap message
// that will be sent in JSON format.
type CreateUserResponse struct {
	ID string `json:"id"`
}

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

// GetUserByID will handle user search.
func (uh UserHandler) GetUserByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		id := mux.Vars(req)["id"]

		if err := validateID(id); err != nil {
			uh.logger.Warnf("Invalid ID in the request: %s", id)
			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

			return
		}

		entUser, err := uh.usecase.GetUser(id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				uh.logger.Debugw("No user found.", "ID", id)
				WriteJSONResponse(w, http.StatusNotFound, Response{Message: MsgUserNotFound}, uh.logger)

				return
			}

			uh.logger.Errorw("Internal error while searching user.", "ID", id, "error", err.Error())
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, uh.logger)

			return
		}

		user := User{
			ID:        entUser.ID,
			FirstName: entUser.FirstName,
			LastName:  entUser.LastName,
			Email:     entUser.Email,
			CreatedAt: entUser.CreatedAt,
		}

		uh.logger.Debugw("User successfully found", "ID", id)
		WriteJSONResponse(w, http.StatusOK, user, uh.logger)
	})
}
