package errors

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

type ErrIdentifierNotFound struct {
	node ast.Node
	name string
}

func NewErrIdentifierNotFound(node ast.Node, name string) ErrIdentifierNotFound {
	return ErrIdentifierNotFound{node: node, name: name}
}

func (e ErrIdentifierNotFound) Error() string {
	return fmt.Sprintf("identifier not found %s", e.name)
}
