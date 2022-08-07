package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewNamedExpressionList(id int64, exprs map[string]Expr) *NamedExpressionList {
	return &NamedExpressionList{
		id:    id,
		exprs: exprs,
	}
}

type NamedExpressionList struct {
	id    int64
	exprs map[string]Expr
}

func (el *NamedExpressionList) ID() int64            { return el.id }
func (el *NamedExpressionList) NodeKey() ast.NodeKey { return ast.KeyNamedExpressionList }

func (el *NamedExpressionList) Check(
	env *environment.Environment,
	validMngr validationManager,
) (*object.NamedResult, *execAst.NamedExpressionList, error) {
	res := object.NewNamedResult()
	exprAstMap := map[string]execAst.Expr{}
	for name, expr := range el.exprs {
		exprRes, exprAst, err := expr.Check(env, validMngr)
		if err != nil {
			return nil, nil, err
		}
		exprAstMap[name] = exprAst
		res.ObjectList[name] = exprRes.ObjectList[0]
	}
	return res, execAst.NewNamedExpressionList(el.id, exprAstMap), nil
}
