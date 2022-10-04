// Package entities will consist of key entities  of our project
// and represents core/domain layer of our app -  correct me if I am wrong
// and would not have any dependencies on other layers.
// ent package also contains validations and custom errors.
package entities

import (
	"errors"
	"strconv"
)

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
	if err := checkName(u.FirstName); err != nil {
		return err
	}
	if u.LastName != "" {
		if err := checkName(u.LastName); err != nil {
			return err
		}
	}

	if err := checkMail(u.Email); err != nil {
		return err
	}

	if err := checkPass(u.Password); err != nil {
		return err
	}

	return nil
}

func checkName(name string) error {
	const (
		minNameLen         int    = 2   // 2
		maxNameLen         int    = 256 // 256
		invalidNameSymbols string = "!\"#$%&*+,./:;<=>?@[\\]^_`{|}~1234567890"
	)
	// Checking the length
	nameLen := len(name)
	switch {
	case nameLen < minNameLen:
		return errors.New("Name is too short.")
	case nameLen > maxNameLen:
		return errors.New("Name is too long.")
	}
	// Checking for not allowed symbols
	for _, v := range name {
		for _, e := range invalidNameSymbols {
			if v == e {
				return errors.New("Name contains forbidden symbol: " + string(v) + ".")
			}
		}
	}
	return nil
}

func checkMail(email string) error {
	const (
		maxLocalEmailLen    int    = 64  // 64
		maxDomainEmailLen   int    = 255 // 255
		validEmailSymbols   string = "!#$%&'*+-/=?^_`{|}~"
		invalidEmailSymbols string = "\"(),:;<>[\\]"
	)
	// Checking the total length (allowed no more than 64+1+255=320 symbols)
	mailLen := len(email)
	if mailLen > maxLocalEmailLen+1+maxDomainEmailLen {
		return errors.New("Email contents too many symbols: " + strconv.Itoa(mailLen))
	}
	// Check if "at" is not first or last symbol, and it's number is 1
	var atPosition, dotPosition, atNum, dotNum int
	for i, v := range email {
		if v == '@' {
			atPosition = i + 1
			atNum += 1
			if atPosition == 1 || atPosition == mailLen {
				return errors.New("Email can not to begin or to end with @.")
			}
		}
		if v == '.' {
			dotPosition = i + 1
			dotNum += 1 //Checking dots quantity
			if dotPosition == 1 || dotPosition == mailLen {
				return errors.New("Email can not to begin or to end with dot.")
			}
		}
		if v < '!' || v > '~' { // Checking for non-ASCII symbols
			return errors.New("Email contains non-latin or forbidden symbol: " + string(v) + ".")
		}
		for _, e := range invalidEmailSymbols { // Check not allowed symbols
			if v == e {
				return errors.New("Email contains forbidden symbol: " + string(v) + ".")
			}
		}
	}
	switch {
	case dotPosition < atPosition:
		return errors.New("Email does not have dot in the domain part (after @).")
	case atNum != 1:
		return errors.New("Email does not have @ or includes more than 1.")
	case dotNum == 0:
		return errors.New("Email does not have dots.")
	case atPosition > maxLocalEmailLen+1:
		return errors.New("Email have too long local part (before @).")
	case mailLen-atPosition > maxDomainEmailLen:
		return errors.New("Email have too long domain part (after @).")
	}
	return nil
}

func checkPass(pass string) error {
	const (
		minPassLen int = 8   // 8
		maxPassLen int = 256 // 256
	)
	// Checking the length
	passLen := len(pass)
	if passLen < minPassLen {
		return errors.New("Password must have at least 8 characters.")
	}
	if passLen > maxPassLen {
		return errors.New("Password can not be more than 256 characters.")
	}
	// Checking for non-ASCII symbols
	for _, v := range pass {
		if v < '!' || v > '~' {
			return errors.New("Password can contain only Aa-Zz letters, 0-9 digits, and symbols !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
		}
	}
	return nil
}
