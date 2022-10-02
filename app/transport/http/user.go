package http

import (
	"net/http"

	"github.com/go-devs-ua/octagon/app/entities"
)

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

func (uh UserHandler) Map() []HandlerMap {
	return []HandlerMap{
		{EndPoint: "/user", Handler: uh.CreateUser(), Method: http.MethodGet},
		// {EndPoint: "/user", Handler: uh.DeleteUser(), Method: http.MethodDelete},
		// and so on...
	}
}

// CreateUser will handle user creation
func (uh UserHandler) CreateUser() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		usr := entities.User{}

		if err := usr.Validate(); err != nil {
			return
		}

		if err := uh.usecase.Signup(usr); err != nil {
			return
		}
	})
}
