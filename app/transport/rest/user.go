package rest

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-devs-ua/octagon/app/achtung"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUserResponse will wrap message
// that will be sent in JSON format.
type CreateUserResponse struct {
	ID string `json:"id"`
}

// ListPublicUsers holds on all public users ready to present.
type ListPublicUsers struct {
	Results []*entities.PublicUser `json:"results"`
	Next    *url.URL               `json:"next,omitempty"`
}

type QueryParams map[string]any

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

			if errors.Is(err, achtung.ErrDuplicateEmail) {
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
// `limit` specified in request, if none was specified, defaultLimit is pagination limit default.
// `next` link is present if only `offset` and `limit` is less than total number of objects.
// defaultSort will be used if no params were specified.
func (uh UserHandler) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var (
			users       []*entities.PublicUser
			queryParams = QueryParams{offset: defaultOffset, limit: defaultLimit, sort: defaultSort}
			err         error
			val         any
			numUsers    int
			next        *url.URL
			query       url.Values
		)

		for param, arg := range req.URL.Query() {
			if param == offset || param == limit {
				val, err = strconv.Atoi(arg[0])
				if err != nil {
					WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: err.Error()}, uh.logger)
					uh.logger.Errorf("Failed parsing query params: %v+", err)
				}
			}
			queryParams[param] = val
		}

		ctx, cancel := context.WithTimeout(req.Context(), queryTimeoutSeconds*time.Second)
		defer cancel()

		users, err = uh.usecase.Fetch(ctx, queryParams)
		if err != nil {
			uh.logger.Errorf("Failed fetching users from repository: %v", err)
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not get all users"}, uh.logger)

			return
		}

		// TODO: Pagination.
		_, _ = numUsers, query
		//numUsers = len(users)
		//if numUsers == 0 {
		//	WriteJSONResponse(w, http.StatusNoContent, Response{Message: MsgNoContent}, uh.logger)
		//
		//	return
		//}
		//
		//nextIs := (queryParams[offset].(int) + queryParams[limit].(int)) < numUsers
		//
		//if nextIs {
		//	next = req.URL
		//	query.Set(offset, offset+limit)
		//	query.Set(limit, limit+limit)
		//	next.RawQuery = query.Encode()
		//}

		//w.Header().Set("Content-Range", fmt.Sprintf("objects %v-%v/%d", offset, limit, numUsers))

		WriteJSONResponse(w, http.StatusOK, ListPublicUsers{Results: users, Next: next}, uh.logger)
	})
}
