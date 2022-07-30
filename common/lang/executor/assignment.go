package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
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

func (a *Assignment) Exec(env *fdalang.Environment, result *Result, execQueue *ExecFnList) error {
	fn := execQueue.AddAfterCurrent(a.value, func() error {
		return a.value.Exec(env, result, execQueue)
	})
	for i := range a.left {
		ii := i
		fn = execQueue.AddAfter(fn, a.left[ii], func() error {
			varName := a.left[ii].value
			env.Set(varName, result.objectList[ii])
			return nil
		})
	}
	return nil
}
