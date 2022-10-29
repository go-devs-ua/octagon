package usecase

import (
	"errors"
	"fmt"

	"github.com/go-devs-ua/octagon/app/entities"
	"github.com/go-devs-ua/octagon/app/globals"
)

type User struct {
	Repo UserRepository
}

// NewUser is a famous  trick with accepting
// interfaces and returning struct.
func NewUser(repo UserRepository) User {
	return User{Repo: repo}
}

// SignUp represents business logic
// and will take care of creating user.
func (u User) SignUp(user entities.User) (string, error) {
	id, err := u.Repo.AddUser(user)
	if err != nil {
		if errors.Is(err, globals.ErrDuplicateEmail) {
			return "", globals.ErrDuplicateEmail
		}

		return "", fmt.Errorf("error while adding user to database: %w", err)
	}

	return id, nil
}

// Get takes care of finding user by ID.
func (u User) Get(id string) (*entities.User, error) {
	user, err := u.Repo.FindUser(id)
	if err != nil {
		return nil, fmt.Errorf("error while searching user in database: %w", err)
	}

	return user, nil
}

// GetAll retrieves all suitable users from repository.
func (u User) GetAll(offset, limit, sort string) ([]entities.User, error) {
	users, err := u.Repo.GetUsers(offset, limit, sort)
	if err != nil {
		return nil, fmt.Errorf("error fetching users: %w", err)
	}

	return users, nil
}
