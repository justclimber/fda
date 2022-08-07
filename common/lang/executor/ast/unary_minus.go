package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewUnaryMinus(value Expr) *UnaryMinus {
	return &UnaryMinus{
		value: value,
	}
}

type UnaryMinus struct {
	id    int64
	value Expr
}

func (um *UnaryMinus) ID() int64            { return um.id }
func (um *UnaryMinus) NodeKey() ast.NodeKey { return ast.KeyUnaryMinus }

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
