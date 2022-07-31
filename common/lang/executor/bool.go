package executor

func NewBool(value bool) *Bool {
	return &Bool{
		key:   KeyBool,
		value: value,
	}
}

type Bool struct {
	id    int64
	key   NodeKey
	value bool
}

func (b *Bool) NodeKey() NodeKey { return b.key }
func (b *Bool) ID() int64        { return b.id }

func (b *Bool) Exec(_ *Environment, result *Result, executor execManager) error {
	executor.AddNextExec(b, func() error {
		result.Add(toReservedBoolObj(b.value))
		return nil
	})
	return nil
}
