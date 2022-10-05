package entities

import (
	"errors"
	"regexp"
	"strconv"
)

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

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
		minNameLen int = 2
		maxNameLen int = 256
	)
	// Checking the length
	nameLen := len(name)
	switch {
	case nameLen < minNameLen:
		return errors.New("name is too short")
	case nameLen > maxNameLen:
		return errors.New("name is too long")
	}
	// Checking for forbidden symbols
	if valid, _ := regexp.MatchString(`^[a-zA-Zа-яА-ЯҐІЇЄґіїє'\s]{2,256}$`, name); !valid {
		return errors.New("invalid name")
	}

	return nil
}

func checkMail(email string) error {
	const (
		maxLocalEmailLen  int = 64
		maxDomainEmailLen int = 255
	)
	// Checking the total length (allowed no more than 64+1+255=320 symbols)
	mailLen := len(email)
	if mailLen > maxLocalEmailLen+1+maxDomainEmailLen {
		return errors.New("email contents too many symbols: " + strconv.Itoa(mailLen))
	}
	// Checking for some email issues by regular expression
	if valid, _ := regexp.MatchString(`^([^\.@(),:;<>@[\\\]][\da-z!#$%&'*+-/=?^_\x60{|}~]+)@([^\-][\da-z-\.]+[^\-]+)+\.([a-z]{2,6})$`, email); !valid {
		return errors.New("invalid email")
	}

	return nil
}

func checkPass(pass string) error {
	const (
		minPassLen int = 8
		maxPassLen int = 256
	)
	// Checking the length
	passLen := len(pass)
	if passLen < minPassLen {
		return errors.New("password must have at least 8 characters")
	}
	if passLen > maxPassLen {
		return errors.New("password can not be more than 256 characters")
	}
	// Checking for non-ASCII symbols
	if valid, _ := regexp.MatchString(`^[[:graph:]]$`, pass); !valid {
		return errors.New("invalid password")
	}

	return nil
}
