package entities

import (
	"fmt"
	"regexp"
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

const (
	nameMask = `^[\p{L}&\s-\\'’.]{2,256}$`
	mailMask = `(?i)^(?:[a-z\d!#$%&'*+/=?^_\x60{|}~-]+(?:\.[a-z\d!#$%&'*+/=?^_\x60{|}~-]+)*)@(?:(?:[a-z\d](?:[a-z\d-]*[a-z\d])?\.)+[a-z\d](?:[a-z\d-]*[a-z\d])?)$`
	passMask = `^[[:graph:]]{8,256}$`
)

var (
	nameRegex = regexp.MustCompile(nameMask)
	mailRegex = regexp.MustCompile(mailMask)
	passRegex = regexp.MustCompile(passMask)
)

func (u User) Validate() error {
	if err := checkName(u.FirstName); err != nil {
		return fmt.Errorf("invalid first name: %v", err.Error())
	}

	if len(u.LastName) > 1 {
		if err := checkName(u.LastName); err != nil {
			return fmt.Errorf("invalid last name: %v", err.Error())
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
	if valid := nameRegex.MatchString(name); !valid {
		return fmt.Errorf("name does not match with regex: `%s`", nameMask)
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
		return fmt.Errorf("local part of email contains too many bytes: %v", atIndex)
	}
	if localPartLen := len(email) - atIndex - 1; localPartLen > maxDomainBytes {
		return fmt.Errorf("domain part of email contains too many bytes: %v", localPartLen)
	}
	// Checking for other email issues by regular expression
	if valid := mailRegex.MatchString(email); !valid {
		return fmt.Errorf("email does not match with regex: `%s`", mailMask)
	}

	return nil
}

func checkPass(password string) error {
	if valid := passRegex.MatchString(password); !valid {
		return fmt.Errorf("password does not match with regex: `%s`", passMask)
	}

	return nil
}

// Hide sensitive data while printing
func (u User) String() string {
	return fmt.Sprintf(`{"ID":"%v"}`, u.ID)
}
