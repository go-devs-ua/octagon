package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-devs-ua/octagon/app/repository/pg"

	"github.com/go-devs-ua/octagon/app/entities"
)

type User struct {
	Repo UserRepository
}

// NewUser is a famous  trick with accepting
// interfaces and returning structs.
func NewUser(repo UserRepository) User {
	return User{Repo: repo}
}

// Signup represents business logic
// and will take care of creating user.
func (u User) Signup(user entities.User) (string, error) {
	id, err := u.Repo.Add(user)
	if err != nil {
		if errors.Is(err, pg.ErrDuplicateEmail) {
			return "", ErrDuplicateEmail
		}

		return "", fmt.Errorf("error while adding user to database: %w", err)
	}

	return id, nil
}

// Fetch retrieves all suitable users from repository.
func (u User) Fetch(ctx context.Context, params map[string]any) ([]*entities.PublicUser, error) {
	users, err := u.Repo.GetAll(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}

	return users, nil
}
