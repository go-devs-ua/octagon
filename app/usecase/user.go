package usecase

import (
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
	// TODO: Some magic
	ID, err := u.Repo.Add(user)
	if err != nil {
		return "", err
	}

	return ID, nil
}
