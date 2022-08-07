package errors

import (
	"fmt"
	"strings"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrorType int

const (
	ErrorTypeUnknown ErrorType = iota + 1
	ErrorTypeIdentifierNotFound
	ErrorTypeMismatch
)

func (et ErrorType) String() string {
	return errorMessages[et]
}

var errorMessages = [...]string{
	ErrorTypeUnknown:            "unknown error occurred",
	ErrorTypeIdentifierNotFound: "identifier not found",
}

type ValidationError struct {
	node    ast.Node
	errType ErrorType
}

func NewValidationError(node ast.Node, errType ErrorType) *ValidationError {
	return &ValidationError{
		errType: errType,
		node:    node,
	}
}

func (r *ValidationError) Error() string {
	return fmt.Sprintf("%s\n", r.errType.String())
}

func NewValidationErrorSet() *ValidationErrorSet {
	return &ValidationErrorSet{
		errorSet: make([]*ValidationError, 0),
	}
}

type ValidationErrorSet struct {
	errorSet []*ValidationError
}

func (vs *ValidationErrorSet) Add(e *ValidationError) {
	vs.errorSet = append(vs.errorSet, e)
}

func (vs *ValidationErrorSet) Empty() bool {
	return len(vs.errorSet) == 0
}

func (vs *ValidationErrorSet) Error() string {
	var s []string
	for _, validationError := range vs.errorSet {
		s = append(s, validationError.Error())
	}
	return strings.Join(s, "\n")
}

type RuntimeError struct {
	node    ast.Node
	errType ErrorType
}

func NewRuntimeError(node ast.Node, errType ErrorType) *RuntimeError {
	return &RuntimeError{
		errType: errType,
		node:    node,
	}
}

func (r *RuntimeError) Error() string {
	return fmt.Sprintf("%s\n", r.errType.String())
}
