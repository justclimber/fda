package ast

func NewStructDefinition(name string, fields []*VarAndType) *StructDefinition {
	return &StructDefinition{
		key:    KeyStructDefinition,
		name:   name,
		Fields: fields,
	}
}

type StructDefinition struct {
	id     int64
	key    NodeKey
	name   string
	Fields []*VarAndType
}

func (sd *StructDefinition) ID() int64        { return sd.id }
func (sd *StructDefinition) NodeKey() NodeKey { return sd.key }
