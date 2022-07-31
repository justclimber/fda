package executor

func NewStructDefinition(name string, fields []*VarAndType) *StructDefinition {
	return &StructDefinition{
		key:    KeyStructDefinition,
		name:   name,
		fields: fields,
	}
}

type StructDefinition struct {
	id     int64
	key    NodeKey
	name   string
	fields []*VarAndType
}

func (sd *StructDefinition) ID() int64        { return sd.id }
func (sd *StructDefinition) NodeKey() NodeKey { return sd.key }
