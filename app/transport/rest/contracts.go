package rest

import (
	"github.com/go-devs-ua/octagon/app/entities"
)

// UserUsecase represents User use-case layer.
type UserUsecase interface {
	SignUp(entities.User) (string, error)
	Fetch(string, string, string) ([]entities.User, error)
	Get(id string) (*entities.User, error)
}
