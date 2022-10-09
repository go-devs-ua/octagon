package entities

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
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
		maxLocalBytes  int = 64
		maxDomainBytes int = 255
	)
	// Checking the lengths of local and domain parts
	atIndex := strings.IndexByte(email, '@')
	if atIndex > maxLocalBytes {
		return errors.New("local part of email contains too many bytes: " + strconv.Itoa(atIndex))
	}
	if localPartLen := len(email) - atIndex - 1; localPartLen > maxDomainBytes {
		return errors.New("domain part of email contains too many bytes: " + strconv.Itoa(localPartLen))
	}
	// Checking for other email issues by regular expression
	if valid, _ := regexp.MatchString("(?i)"+`^(?:[a-z\d!#$%&'*+/=?^_\x60{|}~-]+(?:\.[a-z\d!#$%&'*+/=?^_\x60{|}~-]+)*)@(?:(?:[a-z\d](?:[a-z\d-]*[a-z\d])?\.)+[a-z\d](?:[a-z\d-]*[a-z\d])?)$`, email); !valid {
		return errors.New("invalid email (does not match ^(?:[a-z\\d!#$%&'*+/=?^_\\x60{|}~-]+(?:\\.[a-z\\d!#$%&'*+/=?^_\\x60{|}~-]+)*)@(?:(?:[a-z\\d](?:[a-z\\d-]*[a-z\\d])?\\.)+[a-z\\d](?:[a-z\\d-]*[a-z\\d])?)$")
	}

	return nil
}

func checkPass(password string) error {
	if valid, _ := regexp.MatchString(`^[[:graph:]]{8,256}$`, password); !valid {
		return errors.New("invalid password")
	}

	return nil
}
