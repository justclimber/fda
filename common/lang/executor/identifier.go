package executor

func NewIdentifierList(values []string) []*Identifier {
	result := make([]*Identifier, 0, len(values))
	for _, value := range values {
		result = append(result, NewIdentifier(value))
	}

	return result
}

func NewIdentifier(value string) *Identifier {
	return &Identifier{
		key:   KeyIdentifier,
		value: value,
	}
}

type Identifier struct {
	id    int64
	key   NodeKey
	value string
}

func (i *Identifier) ID() int64        { return i.id }
func (i *Identifier) NodeKey() NodeKey { return i.key }

func (i *Identifier) Exec(env *Environment, result *Result, executor execManager) error {
	if val, ok := env.Get(i.value); ok {
		result.objectList[0] = val
		return nil
	}

	return NewRuntimeError(i, ErrorTypeIdentifierNotFound)
}
