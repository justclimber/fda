package executor

import (
	"github.com/justclimber/fda/common/lang/fdalang"
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

func (v *VoidedExpression) Exec(env *fdalang.Environment, execQueue *ExecFnList) error {
	execQueue.AddAfterCurrent(v.expr, func() error {
		return v.expr.Exec(env, NewResult(), execQueue)
	})
	return nil
}
