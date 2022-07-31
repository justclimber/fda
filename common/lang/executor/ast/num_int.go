package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
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

func (n *NumInt) Exec(_ *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(n, func() error {
		result.Add(&object.ObjInteger{Value: n.value})
		return nil
	})
	return nil
}
