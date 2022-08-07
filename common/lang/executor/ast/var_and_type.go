package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewVarAndType(varName string, varType object.ObjectType) *VarAndType {
	return &VarAndType{
		varName: varName,
		varType: varType,
	}
}

type VarAndType struct {
	id      int64
	varName string
	varType object.ObjectType
}

func (vt *VarAndType) ID() int64            { return vt.id }
func (vt *VarAndType) NodeKey() ast.NodeKey { return ast.KeyVarAndType }
