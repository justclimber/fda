package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/validator/result"
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

func (el *ExpressionList) Check(env ValidatorEnv, validMngr validationManager) (*result.Result, execAst.Expr, error) {
	res := result.NewResult()
	errContainer := errors.NewErrContainer(el)
	exprListAst := make([]execAst.Expr, 0, len(el.exprs))
	for i := range el.exprs {
		r, exprAst, err := el.exprs[i].Check(env, validMngr)
		if err != nil {
			errContainer.Add(err)
			continue
		}
		exprListAst = append(exprListAst, exprAst)
		res.Merge(r)
	}
	if errContainer.NotEmpty() {
		return nil, nil, errContainer
	}

	return res, execAst.NewExpressionList(el.id, exprListAst), nil
}
