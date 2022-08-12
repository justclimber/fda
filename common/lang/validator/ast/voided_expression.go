package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
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
func (v *VoidedExpression) NodeKey() ast.NodeKey { return ast.KeyStatementsBlock }

func (v *VoidedExpression) Check(env ValidatorEnv, validMngr validationManager) (execAst.Stmt, error) {
	_, exprAst, err := v.expr.Check(env, validMngr)
	return execAst.NewVoidedExpression(v.id, exprAst), err
}
