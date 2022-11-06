package rest

import (
	"github.com/go-devs-ua/octagon/app/entities"
)

//go:generate moq -out user_usecase_mock_test.go . UserUsecase

// UserUsecase represents User use-case layer.
type UserUsecase interface {
	SignUp(entities.User) (string, error)
	GetAll(entities.QueryParams) ([]entities.User, error)
	GetByID(id string) (*entities.User, error)
	Delete(entities.User) error
}
