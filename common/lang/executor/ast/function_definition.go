package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewFunctionDefinition(
	name string,
	statementsBlock *StatementsBlock,
	args []*VarAndType,
	returns []*VarAndType,
) *FunctionDefinition {
	return &FunctionDefinition{
		Name:            name,
		statementsBlock: statementsBlock,
		args:            args,
		returns:         returns,
	}
}

type FunctionDefinition struct {
	id              int64
	Name            string
	statementsBlock *StatementsBlock
	args            []*VarAndType
	returns         []*VarAndType
}

func (fd *FunctionDefinition) ID() int64            { return fd.id }
func (fd *FunctionDefinition) NodeKey() ast.NodeKey { return ast.KeyFunctionDefinition }
