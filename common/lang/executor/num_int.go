package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

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

func (n *NumInt) Exec(_ *fdalang.Environment, result *Result, execQueue *ExecFnList) error {
	execQueue.AddAfterCurrent(n, func() error {
		result.Add(&fdalang.ObjInteger{Value: n.value})
		return nil
	})
	return nil
}
