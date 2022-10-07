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
	if valid, _ := regexp.MatchString("(?i)"+`^([^\.@(),:;<>@[\\\]]?[\da-z!#$%&'*+,\-/=?^_\x60{|}~]+.?)+@([^.-][\da-z-]*[^.-]\.?)+([a-z]{2,6})$`, email); !valid {
		return errors.New("email does not match with regexp \"(?i)\"+`^([^\\.@(),:;<>@[\\\\\\]]?[\\da-z!#$%&'*+,\\-/=?^_\\x60{|}~]+.?)+@([^.-][\\da-z-]*[^.-]\\.?)+([a-z]{2,6})$`")
	}

	return nil
}

func checkPass(password string) error {
	if valid, _ := regexp.MatchString(`^[[:graph:]]{2,256}$`, password); !valid {
		return errors.New("invalid password")
	}

	return nil
}
