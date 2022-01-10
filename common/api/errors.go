package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RegisterUserAlreadyExists uint32 = iota + 1
	RegisterUserNameEmpty
	RegisterPasswordEmpty
	LoginUserNotFound
	LoginWrongPassword

	InternalError uint32 = 1000
)

var ErrUnauthorizedInvalidToken = status.Error(codes.Unauthenticated, "auth token is invalid")
