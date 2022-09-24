// Package core will consist of enterprise business rules
// and would not have any dependencies on other layers
// core package also contains validations and custom errors.
package core

// User is key entity in our project
// Entities like User are the least likely to change
// when something external changes.
type User struct {
}

// Using abstract interfaces without specific
// knowledge of their implementation details
// will make our software flexible and maintainable.

// Repository interface can be implemented
// in any kind of repositories like Postgres, MySQL etc.
type Repository interface {
	Add(User) error
}

type Agency interface {
	Signup(User) error
}
