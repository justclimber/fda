package executor

func NewNumInt(value int64) *NumInt {
	return &NumInt{
		key:   KeyNumInt,
		value: value,
	}
}

type NumInt struct {
	id    int64
	key   NodeKey
	value int64
}

func (n *NumInt) ID() int64        { return n.id }
func (n *NumInt) NodeKey() NodeKey { return n.key }

func (n *NumInt) Exec(_ *Environment, result *Result, executor execManager) error {
	executor.AddNextExec(n, func() error {
		result.Add(&ObjInteger{Value: n.value})
		return nil
	})
	return nil
}
