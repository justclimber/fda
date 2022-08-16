package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrFunctionAlreadyExists struct {
	node ast.Node
	name string
}

func NewErrFunctionAlreadyExists(node ast.Node, name string) ErrFunctionAlreadyExists {
	return ErrFunctionAlreadyExists{node: node, name: name}
}

func (e ErrFunctionAlreadyExists) Error() string {
	return fmt.Sprintf("function %s already exists in this package", e.name)
}
