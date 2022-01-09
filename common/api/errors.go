package api

import (
	"errors"
)

const (
	RegisterUserAlreadyExists uint32 = iota + 1
	RegisterUserNameEmpty
	RegisterPasswordEmpty
	LoginUserNotFound
	LoginWrongPassword

	InternalError uint32 = 1000
)

var UnauthorizedError = errors.New("unauthorized access denied")
