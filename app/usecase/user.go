package usecase

import (
	"errors"
	"fmt"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globs"
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
	id, err := u.Repo.AddUser(user)
	if err != nil {
		if errors.Is(err, globs.ErrDuplicateEmail) {
			return "", globs.ErrDuplicateEmail
		}

		return "", fmt.Errorf("error while adding user to database: %w", err)
	}

	return id, nil
}

// GetUser represents business logic
// and will take care of finding user.
func (u User) GetUser(id string) (*entities.User, error) {
	user, err := u.Repo.FindUser(id)
	if err != nil {
		return nil, fmt.Errorf("error while searching user in database: %w", err)
	}

	return user, nil
}
