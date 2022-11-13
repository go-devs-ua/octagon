package rest

import (
	"encoding/json"
	"errors"
	"net/http"

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

func makeUsersRESTful(userArr []entities.User) []User {
	users := make([]User, 0, len(userArr))

	for _, u := range userArr {
		user := User{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			CreatedAt: u.CreatedAt,
		}

		users = append(users, user)
	}

	return users
}

// CreateUser will handle user creation.
func (uh UserHandler) CreateUser(w http.ResponseWriter, req *http.Request) {
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
}

// GetUserByID will handle user search.
func (uh UserHandler) GetUserByID(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	if _, err := uuid.Parse(id); err != nil {
		uh.logger.Warnw("Invalid UUID", "ID", id)
		WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

		return
	}

	user, err := uh.usecase.GetByID(id)
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
}

// GetAllUsers retrieves all entities.User by given parameters.
func (uh UserHandler) GetAllUsers(w http.ResponseWriter, req *http.Request) {
	var params entities.QueryParams

	params.Offset = req.URL.Query().Get("offset")
	params.Limit = req.URL.Query().Get("limit")
	params.Sort = req.URL.Query().Get("sort")

	if err := params.Validate(); err != nil {
		uh.logger.Errorf("Failed validating query: %+v", err)
		WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

		return
	}

	users, err := uh.usecase.GetAll(params)
	if err != nil {
		uh.logger.Errorf("Failed fetching users from repository: %+v", err)
		WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not fetch users"}, uh.logger)

		return
	}

	WriteJSONResponse(w, http.StatusOK, UsersResponse{Results: makeUsersRESTful(users)}, uh.logger)
}

// DeleteUser will handle user creation.
func (uh UserHandler) DeleteUser(w http.ResponseWriter, req *http.Request) {
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

	if err := user.ValidateUUID(); err != nil {
		uh.logger.Warnf("Invalid ID in the request: %s", user)
		WriteJSONResponse(w, http.StatusBadRequest, Response{Message: MsgBadRequest, Details: err.Error()}, uh.logger)

		return
	}

	if err := uh.usecase.Delete(user); err != nil {
		if errors.Is(err, globals.ErrNotFound) {
			uh.logger.Debugw("No user found.", "ID", user.ID)
			WriteJSONResponse(w, http.StatusNotFound, Response{Message: MsgNotFound, Details: err.Error()}, uh.logger)

			return
		}

		uh.logger.Errorw("Internal error while deleting user.", "ID", user.ID, "error", err.Error())
		WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr}, uh.logger)

		return
	}

	WriteJSONResponse(w, http.StatusNoContent, nil, uh.logger)
	uh.logger.Debugw("User successfully deleted", "ID", user.ID)
}
