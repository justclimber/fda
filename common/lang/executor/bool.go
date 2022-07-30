package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

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

func (b *Bool) Exec(env *fdalang.Environment, result *Result, execQueue *ExecFnList) error {
	execQueue.AddAfterCurrent(b, func() error {
		result.Add(toReservedBoolObj(b.value))
		return nil
	})
	return nil
}
