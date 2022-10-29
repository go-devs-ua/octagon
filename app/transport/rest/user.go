package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
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

func makeUsersRESTful(users []entities.User) []User {
	var userArr = make([]User, 0, len(users))

	for _, u := range users {
		user := User{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			CreatedAt: u.CreatedAt,
		}

		userArr = append(userArr, user)
	}

	return userArr
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

		user, err := uh.usecase.Get(id)
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

		userResp := User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
		WriteJSONResponse(w, http.StatusOK, userResp, uh.logger)
	})
}

// QueryParams represent request query params.
type QueryParams struct {
	Offset string `schema:"offset"`
	Limit  string `schema:"limit"`
	Sort   string `schema:"sort"`
}

// GetUsers retrieves all entities.User by given parameters.
func (uh UserHandler) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var params QueryParams

		if err := schema.NewDecoder().Decode(&params, req.URL.Query()); err != nil {
			uh.logger.Errorf("Failed parsing query %+v: %+v", req.URL.Query(), err)
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not parse query"}, uh.logger)

			return
		}

		users, err := uh.usecase.GetAll(params.Offset, params.Limit, params.Sort)
		if err != nil {
			uh.logger.Errorf("Failed fetching users from repository: %+v", err)

			if errors.Is(err, globals.ErrBadQuery) {
				WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: globals.ErrBadQuery.Error()}, uh.logger)

				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not fetch users"}, uh.logger)

			return
		}

		WriteJSONResponse(w, http.StatusOK, UsersResponse{Results: makeUsersRESTful(users)}, uh.logger)
	})
}
