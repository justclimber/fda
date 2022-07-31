package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewVoidedExpression(expr Expr) *VoidedExpression {
	return &VoidedExpression{
		key:  KeyVoidedExpression,
		expr: expr,
	}
}

type VoidedExpression struct {
	id   int64
	key  NodeKey
	expr Expr
}

func (v *VoidedExpression) ID() int64        { return v.id }
func (v *VoidedExpression) NodeKey() NodeKey { return v.key }

func (v *VoidedExpression) Exec(env *environment.Environment, execMngr execManager) error {
	execMngr.AddNextExec(v.expr, func() error {
		return v.expr.Exec(env, object.NewResult(), execMngr)
	})
	return nil
}