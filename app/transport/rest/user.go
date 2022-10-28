package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// CreateUserResponse will wrap message
// that will be sent in JSON format.
type CreateUserResponse struct {
	ID string `json:"id"`
}

// User represents model of entities.User
// specific to transport layer.
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CreatedAt string `json:"created_at"`
}

// UsersResponse holds on array of Users are going to be rendered.
type UsersResponse struct {
	Results []User `json:"results"`
}

func cropUsers(solidUsers []entities.User) []User {
	var users = make([]User, 0, len(solidUsers))

	for _, su := range solidUsers {
		user := User{
			ID:        su.ID,
			Email:     su.Email,
			FirstName: su.FirstName,
			LastName:  su.LastName,
			CreatedAt: su.CreatedAt,
		}
		users = append(users, user)
	}

	return users
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

		id, err := uh.usecase.SignUp(user)
		if err != nil {
			uh.logger.Errorf("Failed creating user: %+v", err)

			if errors.Is(err, globals.ErrDuplicateEmail) {
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

		if _, err := uuid.Parse(id); err != nil {
			uh.logger.Warnw("Invalid UUID", "ID", id)
			WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

			return
		}

		solidUser, err := uh.usecase.Get(id)
		if err != nil {
			if errors.Is(err, globals.ErrNotFound) {
				uh.logger.Debugw("No user found.", "ID", id)
				WriteJSONResponse(w, http.StatusNotFound, Response{Message: MsgNotFound, Details: err.Error()}, uh.logger)

				return
			}
			uh.logger.Errorw("Internal error while searching user.", "ID", id, "error", err.Error())
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, uh.logger)

			return
		}

		user := User{
			ID:        solidUser.ID,
			FirstName: solidUser.FirstName,
			LastName:  solidUser.LastName,
			Email:     solidUser.Email,
			CreatedAt: solidUser.CreatedAt,
		}
		WriteJSONResponse(w, http.StatusOK, user, uh.logger)
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

		users, err := uh.usecase.Fetch(params[offset], params[limit], params[sort])
		if err != nil {
			uh.logger.Errorf("Failed fetching users from repository: %+v", err)
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not fetch users"}, uh.logger)

			return
		}

		WriteJSONResponse(w, http.StatusOK, UsersResponse{Results: cropUsers(users)}, uh.logger)
	})
}
