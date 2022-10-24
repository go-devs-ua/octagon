package usecase

import (
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

// GetUser represents business logic
// and will take care of finding user.
func (u User) GetUser(id string) (entities.PublicUser, error) {
	user, err := u.Repo.Find(id)
	if err != nil {
		if errors.Is(err, pg.ErrInvalidID) {
			return entities.PublicUser{}, ErrInvalidID
		}

		return entities.PublicUser{}, fmt.Errorf("error while searchin user in database: %w", err)
	}

	return user, nil
}
