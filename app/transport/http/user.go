package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/lib/pq"
)

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var (
			user entities.User
			err  error
		)

		defer func() {
			if err != nil {
				uh.logger.LogRequest(req)
			}
		}()

		if err = json.NewDecoder(req.Body).Decode(&user); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{MsgBadRequest}, uh.logger)
			uh.logger.Errorf("Failed decoding %T from JSON: %+v\n", user, err)
			return
		}

		defer func() {
			if err = req.Body.Close(); err != nil {
				uh.logger.Warnf("Failed closing request body: %+v\n", err)
			}
		}()

		if err = user.Validate(); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{"Validation error: " + err.Error()}, uh.logger)
			uh.logger.Errorf("Failed validating %+v: %+v\n", user, err)
			return
		}

		if err = uh.usecase.Signup(user); err != nil {
			uh.logger.Debugf("Failed creating %+v: %+v\n", user, err)

			// TODO: Handle errors gracefully.
			if err, ok := errors.Unwrap(err).(*pq.Error); ok && err.Code.Name() == "unique_violation" {
				WriteJSONResponse(w, http.StatusConflict, Response{MsgEmailConflict}, uh.logger)
				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{MsgInternalSeverErr}, uh.logger)
			return
		}

		WriteJSONResponse(w, http.StatusCreated, Response{MsgUserCreated}, uh.logger)
		uh.logger.Infof("%T successfully created: %+v\n", user, user)
	})
}
