package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

func NewUnaryMinus(value Expr) *UnaryMinus {
	return &UnaryMinus{
		key:   KeyUnaryMinus,
		value: value,
	}
}

type UnaryMinus struct {
	id    int64
	key   NodeKey
	value Expr
}

func (um *UnaryMinus) ID() int64        { return um.id }
func (um *UnaryMinus) NodeKey() NodeKey { return um.key }

func (um *UnaryMinus) Exec(env *fdalang.Environment, result *Result, execQueue *ExecFnList) error {
	fn := execQueue.AddAfterCurrent(um.value, func() error {
		return um.value.Exec(env, result, execQueue)
	})
	execQueue.AddAfter(fn, um, func() error {
		switch obj := result.objectList[0].(type) {
		case *fdalang.ObjInteger:
			obj.Value = -obj.Value
		}
		return nil
	})

	return nil
}
