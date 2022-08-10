package object

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

func NewFunctionDefinition(name string, packageName string, args []*VarAndType, returns []*VarAndType) *FunctionDefinition {
	return &FunctionDefinition{
		Name:    name,
		Package: packageName,
		Args:    args,
		Returns: returns,
	}
}

type FunctionDefinition struct {
	id      int64
	Name    string
	Package string
	Args    []*VarAndType
	Returns []*VarAndType
}

func (fd *FunctionDefinition) ID() int64            { return fd.id }
func (fd *FunctionDefinition) NodeKey() ast.NodeKey { return ast.KeyFunctionDefinition }

func (fd *FunctionDefinition) Type() ObjectType {
	return ObjectType(fmt.Sprintf("%s#%s", fd.Package, fd.Name))
}
