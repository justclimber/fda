package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
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

func (um *UnaryMinus) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(um.value, func() error {
		return um.value.Exec(env, result, execMngr)
	})
	execMngr.AddNextExec(um, func() error {
		switch obj := result.ObjectList[0].(type) {
		case *object.ObjInteger:
			obj.Value = -obj.Value
		}
		return nil
	})

	return nil
}
