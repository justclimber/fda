package ast

import (
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewVarAndType(varName string, varType object.ObjectType) *VarAndType {
	return &VarAndType{
		key:     KeyVarAndType,
		varName: varName,
		varType: varType,
	}
}

type VarAndType struct {
	id      int64
	key     NodeKey
	varName string
	varType object.ObjectType
}

func (vt *VarAndType) NodeKey() NodeKey { return vt.key }
func (vt *VarAndType) ID() int64        { return vt.id }
