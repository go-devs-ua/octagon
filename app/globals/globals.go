package globals

import "errors"

var (
	ErrDuplicateEmail = errors.New("email is already taken")
	ErrNotFound       = errors.New("no user found in DB")
)

const InvalidTextRepresentation = "invalid_text_representation"
