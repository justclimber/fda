package ast

func NewStructDefinition(name string, fields map[string]*VarAndType) *StructDefinition {
	return &StructDefinition{
		key:    KeyStructDefinition,
		Name:   name,
		Fields: fields,
	}
}

type StructDefinition struct {
	id     int64
	key    NodeKey
	Name   string
	Fields map[string]*VarAndType
}

func (sd *StructDefinition) ID() int64        { return sd.id }
func (sd *StructDefinition) NodeKey() NodeKey { return sd.key }
