package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type ErrIfExpressionResultMustBeBool struct {
	node   ast.Node
	actual object.Type
}

func NewErrIfExpressionResultMustBeBool(node ast.Node, actual object.Type) ErrIfExpressionResultMustBeBool {
	return ErrIfExpressionResultMustBeBool{node: node, actual: actual}
}

func (e ErrIfExpressionResultMustBeBool) Error() string {
	return fmt.Sprintf("if expression result must be bool: got %s", e.actual)
}
