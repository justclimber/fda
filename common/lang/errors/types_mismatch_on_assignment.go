package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type ErrTypesMismatchOnAssignment struct {
	node             ast.Node
	existedVarName   string
	expected, actual object.Type
}

func NewErrTypesMismatchOnAssignment(
	node ast.Node,
	existedVarName string,
	expected object.Type,
	actual object.Type,
) ErrTypesMismatchOnAssignment {
	return ErrTypesMismatchOnAssignment{node: node, existedVarName: existedVarName, expected: expected, actual: actual}
}

func (e ErrTypesMismatchOnAssignment) Error() string {
	return fmt.Sprintf(
		"type mismatch on assignment: var %s already have type %s, and new value has type %s",
		e.existedVarName,
		e.expected,
		e.actual,
	)
}
