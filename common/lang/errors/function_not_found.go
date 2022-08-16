package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrFunctionNotFound struct {
	node        ast.Node
	name        string
	packageName string
}

func NewErrFunctionNotFound(node ast.Node, name string, packageName string) ErrFunctionNotFound {
	return ErrFunctionNotFound{node: node, name: name, packageName: packageName}
}

func (e ErrFunctionNotFound) Error() string {
	return fmt.Sprintf("function not found %s in the package %s", e.name, e.packageName)
}
