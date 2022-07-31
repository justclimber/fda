package executor

func NewVarAndType() *VarAndType {
	return &VarAndType{
		key: KeyVarAndType,
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

func (vt *VarAndType) Exec(env *Environment, executor execManager) error {

	return nil
}
