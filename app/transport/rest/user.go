package rest

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

// UserUsecase represents User use-case layer
type UserUsecase interface {
	Signup(entities.User) error
}

// UserHandler is User HTTP handler
// which consist of embedded UserUsecase interface
type UserHandler struct {
	usecase UserUsecase
}

// NewUserHandler will return a new instance
// of UserHandler struct accepting UserUsecase interface
func NewUserHandler(usecase UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

// CreateUser will handle user creation
func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	usr := entities.User{}

	if err := usr.Validate(); err != nil {
		return
	}

	if err := uh.usecase.Signup(usr); err != nil {
		return
	}
}
