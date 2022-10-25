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

// Delete represents business logic
// and will take care of deleting user.
func (u User) Delete(user entities.UserID) error {
	err := u.Repo.Delete(user)
	if err != nil {
		return err
	}

	return nil
}

// Delete represents business logic
// and will take care of deleting user.
func (u User) IsUserExists(user entities.UserID) (bool, error) {
	isExists, err := u.Repo.IsUserExists(user)
	if err != nil {
		return isExists, err
	}

	return isExists, nil
}
