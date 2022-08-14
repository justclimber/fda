package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type ErrTypesMismatch struct {
	node             ast.Node
	expected, actual object.Type
}

func NewErrTypesMismatch(node ast.Node, expected object.Type, actual object.Type) ErrTypesMismatch {
	return ErrTypesMismatch{node: node, expected: expected, actual: actual}
}

func (e ErrTypesMismatch) Error() string {
	return fmt.Sprintf("type mismatch: expected %s, got %s", e.expected, e.actual)
}
