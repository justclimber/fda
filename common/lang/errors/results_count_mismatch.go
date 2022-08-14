package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrResultsCountMismatch struct {
	node             ast.Node
	expected, actual int
}

func NewErrResultsCountMismatch(node ast.Node, expected, actual int) ErrResultsCountMismatch {
	return ErrResultsCountMismatch{node: node, expected: expected, actual: actual}
}

func (e ErrResultsCountMismatch) Error() string {
	return fmt.Sprintf("results count mismatch: expected %d, got %d", e.expected, e.actual)
}
