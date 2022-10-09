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
	if valid, _ := regexp.MatchString(`^[\p{L}&\s-\\'’.]{2,256}$`, name); !valid {
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
	emailLen := len(email)
	if emailLen > maxLocalEmailLen+1+maxDomainEmailLen {
		return errors.New("email contents too many symbols: " + strconv.Itoa(emailLen))
	}
	// Checking for some email issues by regular expression
	if valid, _ := regexp.MatchString("(?i)"+`^(?:[a-z0-9!#$%&'*+/=?^_\x60{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_\x60{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])$`, email); !valid {
		return errors.New("invalid email (does not match ^(?:[a-z0-9!#$%&'*+/=?^_\\x60{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_\\x60{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])$")
	}

	return nil
}

func checkPass(password string) error {
	if valid, _ := regexp.MatchString(`^[[:graph:]]{8,256}$`, password); !valid {
		return errors.New("invalid password")
	}

	return nil
}
