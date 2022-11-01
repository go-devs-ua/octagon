// Package usecase holds on services
// that build all together the business flow,
// and represents so-called business logic layer of the app
// usecase only depends on package ent.
package usecase

import (
	"github.com/go-devs-ua/octagon/app/entities"
)

// Repository interface can be implemented
// in any kind of repositories like Postgres, MySQL etc.
type Repository interface {
	AddUser(entities.User) (string, error)
	FindUser(string) (*entities.User, error)
	GetAllUsers(entities.QueryParams) ([]entities.User, error)
	DeleteUser(entities.User) error
}
