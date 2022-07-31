package executor

func NewStructDefinition(fields *Assignment) *StructDefinition {
	return &StructDefinition{
		key: KeyStructDefinition,
	}
}

type StructDefinition struct {
	id     int64
	key    NodeKey
	name   string
	fields string
}

func (sd *StructDefinition) ID() int64        { return sd.id }
func (sd *StructDefinition) NodeKey() NodeKey { return sd.key }

func (sd *StructDefinition) Exec(env *Environment) error {
	return nil
}
