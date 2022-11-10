package rest

import (
	"github.com/go-devs-ua/octagon/app/entities"
)

//go:generate mockgen -source=./contracts.go -destination=./mock_usecase_test.go -package=rest

// UserUsecase represents User use-case layer.
type UserUsecase interface {
	SignUp(entities.User) (string, error)
	GetAll(entities.QueryParams) ([]entities.User, error)
	GetByID(id string) (*entities.User, error)
	Delete(entities.User) error
}
