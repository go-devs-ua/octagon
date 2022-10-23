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
	id, err := u.Repo.Add(user)
	if err != nil {
		return "", err
	}

	return id, nil
}

// GetUser represents business logic
// and will take care of finding user.
func (u User) GetUser(id string) (entities.PublicUser, error) {
	user, err := u.Repo.Find(id)
	if err != nil {
		return entities.PublicUser{}, err
	}

	return user, nil
}
