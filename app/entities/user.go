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
		minNameLen         int    = 2
		maxNameLen         int    = 256
		invalidNameSymbols string = "!\"#$%&*+,./:;<=>?@[\\]^_`{|}~1234567890"
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
	for _, v := range name {
		for _, e := range invalidNameSymbols {
			if v == e {
				return errors.New("name contains forbidden symbol: " + string(v))
			}
		}
	}

	return nil
}

func checkMail(email string) error {
	const (
		maxLocalEmailLen    int    = 64
		maxDomainEmailLen   int    = 255
		validEmailSymbols   string = "!#$%&'*+-/=?^_`{|}~"
		invalidEmailSymbols string = "\"(),:;<>[\\]"
	)
	// Checking the total length (allowed no more than 64+1+255=320 symbols)
	mailLen := len(email)
	if mailLen > maxLocalEmailLen+1+maxDomainEmailLen {
		return errors.New("email contents too many symbols: " + strconv.Itoa(mailLen))
	}
	// Checking for some email issues by regular expression
	if valid, _ := regexp.MatchString(`^[^\.\\\(\)\@]{1,}@.{1,}[^-]$`, email); !valid {
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
	for _, v := range pass {
		if v < '!' || v > '~' {
			return errors.New("password can contain only Aa-Zz letters, 0-9 digits, and symbols !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
		}
	}

	return nil
}
