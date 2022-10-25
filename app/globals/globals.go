package globals

import "errors"

var (
	ErrInvalidID      = errors.New("invalid ID")
	ErrDuplicateEmail = errors.New("email is already taken")
	ErrNotFound       = errors.New("no user found in DB")
)

const (
	UniqueViolationErrCode    = "unique_violation"
	InvalidTextRepresentation = "invalid_text_representation"
)
