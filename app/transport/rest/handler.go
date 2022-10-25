package rest

import (
	"github.com/go-devs-ua/octagon/lgr"
)

// UserHandler is User HTTP handler
// which consist of embedded Usecase interface.
type UserHandler struct {
	usecase Usecase
	logger  *lgr.Logger
}

// NewUserHandler will return a new instance
// of UserHandler struct accepting Usecase interface.
func NewUserHandler(usecase Usecase, logger *lgr.Logger) UserHandler {
	return UserHandler{
		usecase: usecase,
		logger:  logger,
	}
}
