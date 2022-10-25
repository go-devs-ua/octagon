package rest

import (
	"context"

	"github.com/go-devs-ua/octagon/app/entities"
)

// Usecase represents User use-case layer.
type Usecase interface {
	SignupUser(entities.User) (string, error)
	FetchUsers(context.Context, entities.QueryParams) ([]entities.PublicUser, error)
}
