package usecase

import (
	"github.com/go-devs-ua/octagon/app/ent"
)

type UserLogic struct { // Interactor
	Repo Repository
}

// NewUseCase is a famous
// trick with accepting interfaces and returning structs
func NewUseCase(repo Repository) *UserLogic {
	return &UserLogic{Repo: repo}
}

func (ul *UserLogic) Signup(usr ent.User) error {
	if err := ul.Repo.Add(usr); err != nil {
		return err
	}
	return nil
}
