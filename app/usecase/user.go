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
func (u User) Signup(usr entities.User) error {
	// TODO: Some magic
	if err := u.Repo.Add(usr); err != nil {
		return err
	}

	return nil
}
