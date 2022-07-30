package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
)

func NewIfStatement(condition Expr, trueBranch, falseBranch *StatementsBlock) *IfStatement {
	return &IfStatement{
		key:         KeyIfStatement,
		condition:   condition,
		trueBranch:  trueBranch,
		falseBranch: falseBranch,
	}
}

type IfStatement struct {
	id          int64
	key         NodeKey
	condition   Expr
	trueBranch  *StatementsBlock
	falseBranch *StatementsBlock
}

func (is *IfStatement) NodeKey() NodeKey { return is.key }
func (is *IfStatement) ID() int64        { return is.id }

func (is *IfStatement) Exec(env *fdalang.Environment, execQueue *ExecFnList) error {
	result := NewResult()
	fn := execQueue.AddAfterCurrent(is.condition, func() error {
		return is.condition.Exec(env, result, execQueue)
	})
	execQueue.AddAfter(fn, is.condition, func() error {
		r := result.objectList[0].(*fdalang.ObjBoolean).Value
		if r {
			err := is.trueBranch.Exec(env, execQueue)
			if err != nil {
				return err
			}
		} else if is.falseBranch != nil {
			err := is.falseBranch.Exec(env, execQueue)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}
