package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/validator/result"
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
	env ValidatorEnv,
	validMngr validationManager,
) (*result.NamedResult, *execAst.NamedExpressionList, error) {
	res := result.NewNamedResult()
	errContainer := errors.NewErrContainer(el)
	exprAstMap := map[string]execAst.Expr{}
	for name, expr := range el.exprs {
		exprRes, exprAst, err := expr.Check(env, validMngr)
		errContainer.Add(err)
		if err != nil {
			continue
		}
		if exprRes.Count() != 1 {
			errContainer.Add(errors.NewErrResultsCountMismatch(expr, 1, exprRes.Count()))
		}
		exprAstMap[name] = exprAst
		res.Set(name, exprRes.Get())
	}
	if errContainer.NotEmpty() {
		return nil, nil, errContainer
	}
	return res, execAst.NewNamedExpressionList(el.id, exprAstMap), nil
}
