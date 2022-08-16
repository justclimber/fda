package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrPackageNotFound struct {
	node ast.Node
	name string
}

func NewErrPackageNotFound(node ast.Node, name string) ErrPackageNotFound {
	return ErrPackageNotFound{node: node, name: name}
}

func (e ErrPackageNotFound) Error() string {
	return fmt.Sprintf("package not found %s", e.name)
}
