package ast

func NewVarAndType(varName, varType string) *VarAndType {
	return &VarAndType{
		key:     KeyVarAndType,
		varName: varName,
		varType: varType,
	}
}

type VarAndType struct {
	id      int64
	key     NodeKey
	varType string
	varName string
}

func (vt *VarAndType) NodeKey() NodeKey { return vt.key }
func (vt *VarAndType) ID() int64        { return vt.id }
