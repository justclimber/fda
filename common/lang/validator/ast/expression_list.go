package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewExpressionList(id int64, exprs []Expr) *ExpressionList {
	return &ExpressionList{
		id:    id,
		exprs: exprs,
	}
}

type ExpressionList struct {
	id    int64
	exprs []Expr
}

func (el *ExpressionList) ID() int64            { return el.id }
func (el *ExpressionList) NodeKey() ast.NodeKey { return ast.KeyExpressionList }

func (el *ExpressionList) Check(env *environment.Environment, validMngr validationManager) (*object.Result, execAst.Expr, error) {
	result := object.NewResult()
	exprListAst := make([]execAst.Expr, 0, len(el.exprs))
	for i := range el.exprs {
		r, exprAst, err := el.exprs[i].Check(env, validMngr)
		if err != nil {
			return nil, nil, err
		}
		exprListAst = append(exprListAst, exprAst)
		result.Add(r.ObjectList[0])
	}

	return result, execAst.NewExpressionList(el.id, exprListAst), nil
}
