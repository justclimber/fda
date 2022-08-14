package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrArgCountMismatch struct {
	node             ast.Node
	expected, actual int
}

func NewErrArgCountMismatch(node ast.Node, expected, actual int) ErrArgCountMismatch {
	return ErrArgCountMismatch{node: node, expected: expected, actual: actual}
}

func (e ErrArgCountMismatch) Error() string {
	return fmt.Sprintf("arg count mismatch: expected %d, got %d", e.expected, e.actual)
}
