package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewVoidedExpression(id int64, expr Expr) *VoidedExpression {
	return &VoidedExpression{
		id:   id,
		expr: expr,
	}
}

type VoidedExpression struct {
	id   int64
	expr Expr
}

func (v *VoidedExpression) ID() int64            { return v.id }
func (v *VoidedExpression) NodeKey() ast.NodeKey { return ast.KeyVoidedExpression }

func (v *VoidedExpression) Exec(env *environment.Environment, execMngr execManager) error {
	execMngr.AddNextExec(v.expr, func() error {
		return v.expr.Exec(env, object.NewResult(), execMngr)
	})
	return nil
}
