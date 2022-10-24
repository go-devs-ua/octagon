package usecase

import "errors"

var ErrDuplicateEmail = errors.New("email is already taken")
var ErrInvalidID = errors.New("invalid ID")
