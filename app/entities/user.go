// Package entities will consist of key entities  of our project
// and represents core/domain layer of our app -  correct me if I am wrong
// and would not have any dependencies on other layers.
// ent package also contains validations and custom errors.
package entities

// User is key entity in our project
// Entities like User are the least likely to change
// when something external changes.
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

// TODO: sentinel errors

// Validate will validate User's signup data
func (u User) Validate() error {
	return nil
}
