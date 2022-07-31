package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
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

func (is *IfStatement) Exec(env *environment.Environment, execMngr execManager) error {
	result := object.NewResult()
	execMngr.AddNextExec(is.condition, func() error {
		return is.condition.Exec(env, result, execMngr)
	})
	execMngr.AddNextExec(is.condition, func() error {
		r := result.ObjectList[0].(*object.ObjBoolean).Value
		if r {
			err := is.trueBranch.Exec(env, execMngr)
			if err != nil {
				return err
			}
		} else if is.falseBranch != nil {
			err := is.falseBranch.Exec(env, execMngr)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}
