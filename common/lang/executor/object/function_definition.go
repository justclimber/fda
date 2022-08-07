package object

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewFunctionDefinition(
	name string,
	args []*VarAndType,
	returns []*VarAndType,
) *FunctionDefinition {
	return &FunctionDefinition{
		Name:    name,
		Args:    args,
		Returns: returns,
	}
}

type FunctionDefinition struct {
	id      int64
	Name    string
	Args    []*VarAndType
	Returns []*VarAndType
}

func (fd *FunctionDefinition) ID() int64            { return fd.id }
func (fd *FunctionDefinition) NodeKey() ast.NodeKey { return ast.KeyFunctionDefinition }
