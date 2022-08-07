package object

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewVarAndType(varName string, varType ObjectType) *VarAndType {
	return &VarAndType{
		VarName: varName,
		VarType: varType,
	}
}

type VarAndType struct {
	id      int64
	VarName string
	VarType ObjectType
}

func (vt *VarAndType) ID() int64            { return vt.id }
func (vt *VarAndType) NodeKey() ast.NodeKey { return ast.KeyVarAndType }
