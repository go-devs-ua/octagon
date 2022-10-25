package rest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-devs-ua/octagon/app/alerts"

	"github.com/go-devs-ua/octagon/app/entities"
)

// CreateUserResponse will wrap message
// that will be sent in JSON format.
type CreateUserResponse struct {
	ID string `json:"id"`
}

// ListPublicUsers holds on all public users ready to present.
type ListPublicUsers struct {
	Results []entities.PublicUser `json:"results"`
	Next    string                `json:"next,omitempty"`
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

		id, err := uh.usecase.SignupUser(user)
		if err != nil {
			uh.logger.Errorf("Failed creating user: %+v", err)

			if errors.Is(err, alerts.ErrDuplicateEmail) {
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
		var (
			users  []entities.PublicUser
			params = entities.QueryParams{Sort: defaultSort}
			next   string
			err    error
		)

		for param, arg := range req.URL.Query() {
			val := strings.Join(arg, "")

			switch param {
			case offset:
				params.Offset, err = strconv.Atoi(val)
			case limit:
				params.Limit, err = strconv.Atoi(val)
			case sort:
				params.Sort = val
			}

			if err != nil {
				uh.logger.Errorw("Failed parsing query argument.",
					MsgParam, param,
					MsgArg, val,
					MsgErr, err,
				)
				WriteJSONResponse(w, http.StatusInternalServerError,
					Response{Message: MsgInternalSeverErr, Details: "error parsing query argument: " + val}, uh.logger)

				return
			}
		}

		if params.Limit > maxLimit {
			u := req.URL
			q := &url.Values{}
			q.Set(offset, strconv.Itoa(params.Offset+maxLimit))
			q.Set(limit, strconv.Itoa(params.Limit-maxLimit))
			u.RawQuery = q.Encode()
			next = u.String()

			params.Limit = maxLimit
		}

		ctx, cancel := context.WithTimeout(req.Context(), queryTimeoutSeconds*time.Second)
		defer cancel()

		users, err = uh.usecase.FetchUsers(ctx, params)
		if err != nil {
			uh.logger.Errorf("Failed fetching users from repository: %v", err)
			WriteJSONResponse(w, http.StatusInternalServerError, Response{Message: MsgInternalSeverErr, Details: "could not get all users"}, uh.logger)

			return
		}

		if len(users) == 0 {
			WriteJSONResponse(w, http.StatusNoContent, Response{Message: MsgNoContent}, uh.logger)

			return
		}

		WriteJSONResponse(w, http.StatusOK, ListPublicUsers{Results: users, Next: next}, uh.logger)
	})
}
