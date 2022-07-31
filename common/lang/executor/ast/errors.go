package ast

import (
	"fmt"
)

type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota + 1
	ErrorTypeIdentifierNotFound
)

func (et ErrorType) String() string {
	return errorMessages[et]
}

var errorMessages = [...]string{
	ErrorTypeUnknown:            "unknown error occurred",
	ErrorTypeIdentifierNotFound: "identifier not found",
}

type RuntimeError struct {
	node    Node
	errType ErrorType
}

func NewRuntimeError(node Node, errType ErrorType) *RuntimeError {
	return &RuntimeError{
		errType: errType,
		node:    node,
	}
}

func (r *RuntimeError) Error() string {
	return fmt.Sprintf("%s\n", r.errType.String())
}
