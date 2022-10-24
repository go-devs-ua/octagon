package rest

import (
	"github.com/go-devs-ua/octagon/lgr"
)

// UserHandler is User HTTP handler
// which consist of embedded UserUsecase interface.
type UserHandler struct {
	usecase UserUsecase
	logger  *lgr.Logger
}

// NewUserHandler will return a new instance
// of UserHandler struct accepting UserUsecase interface.
func NewUserHandler(usecase UserUsecase, logger *lgr.Logger) UserHandler {
	return UserHandler{
		usecase: usecase,
		logger:  logger,
	}
}
