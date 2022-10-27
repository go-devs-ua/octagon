// Package usecase holds on services
// that build all together the business flow,
// and represents so-called business logic layer of the app
// usecase only depends on package ent.
package usecase

import (
	"github.com/go-devs-ua/octagon/app/entities"
)

// UserRepository interface can be implemented
// in any kind of repositories like Postgres, MySQL etc.
type UserRepository interface {
	Add(entities.User) (string, error)
	GetUsers(string, string, string) ([]entities.User, error)
}
