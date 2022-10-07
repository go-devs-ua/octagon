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
		user := entities.User{}

		if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{MsgBadRequest}, uh.logger)
			uh.logger.Errorf("%+v\n", err)
			return
		}

		defer func() {
			if err := req.Body.Close(); err != nil {
				uh.logger.Warnf("%+v\n", err)
			}
		}()

		if err := user.Validate(); err != nil {
			WriteJSONResponse(w, http.StatusBadRequest, Response{"Validation error: " + err.Error()}, uh.logger)
			uh.logger.Errorf("%+v\n", err)
			return
		}

		if err := uh.usecase.Signup(user); err != nil {
			uh.logger.Debugf("%+v\n", err)

			// TODO: Handle errors gracefully.
			if err, ok := errors.Unwrap(err).(*pq.Error); ok && err.Code.Name() == "unique_violation" {
				WriteJSONResponse(w, http.StatusConflict, Response{MsgEmailConflict}, uh.logger)
				return
			}

			WriteJSONResponse(w, http.StatusInternalServerError, Response{MsgInternalSeverErr}, uh.logger)
			return
		}

		WriteJSONResponse(w, http.StatusCreated, Response{MsgUserCreated}, uh.logger)
		uh.logger.Infof("%s: %+v\n", MsgUserCreated, user)
	})
}
