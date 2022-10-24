package rest

import (
	"context"
	"github.com/go-devs-ua/octagon/app/entities"
)

// UserUsecase represents User use-case layer.
type UserUsecase interface {
	Signup(entities.User) (string, error)
	Fetch(context.Context, map[string]any) ([]*entities.PublicUser, error)
}
