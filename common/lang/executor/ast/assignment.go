package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewAssignment(left []*Identifier, value Expr) *Assignment {
	return &Assignment{
		key:   KeyAssignment,
		left:  left,
		value: value,
	}
}

type Assignment struct {
	id    int64
	key   NodeKey
	left  []*Identifier
	value Expr
}

func (a *Assignment) NodeKey() NodeKey { return a.key }
func (a *Assignment) ID() int64        { return a.id }

func (a *Assignment) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(a.value, func() error {
		return a.value.Exec(env, result, execMngr)
	})
	for i := range a.left {
		ii := i
		execMngr.AddNextExec(a.left[ii], func() error {
			varName := a.left[ii].value
			env.Set(varName, result.ObjectList[ii])
			return nil
		})
	}
	return nil
}
