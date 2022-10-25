package usecase

import (
	"context"
	"fmt"

	"github.com/go-devs-ua/octagon/app/entities"
)

type User struct {
	Repo Repository
}

// NewUser is a famous  trick with accepting
// interfaces and returning structs.
func NewUser(repo Repository) User {
	return User{Repo: repo}
}

// SignupUser represents business logic
// and will take care of creating user.
func (u User) SignupUser(user entities.User) (string, error) {
	id, err := u.Repo.AddUser(user)
	if err != nil {
		return "", fmt.Errorf("error while adding user to database: %w", err)
	}

	return id, nil
}

// FetchUsers retrieves all suitable users from repository.
func (u User) FetchUsers(ctx context.Context, params entities.QueryParams) ([]entities.PublicUser, error) {
	users, err := u.Repo.GetAllUsers(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}

	return users, nil
}
