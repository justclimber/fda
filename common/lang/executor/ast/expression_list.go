package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewExpressionList(exprs []Expr) *ExpressionList {
	return &ExpressionList{
		exprs: exprs,
	}
}

type ExpressionList struct {
	id    int64
	exprs []Expr
}

func (el *ExpressionList) ID() int64        { return el.id }
func (el *ExpressionList) NodeKey() NodeKey { return KeyExpressionList }

func (el *ExpressionList) Exec(env *environment.Environment, result *object.Result, execMngr execManager) error {
	results := make([]*object.Result, len(el.exprs))
	for i := range el.exprs {
		ii := i
		results[ii] = object.NewResult()
		execMngr.AddNextExec(el.exprs[ii], func() error {
			return el.exprs[ii].Exec(env, results[ii], execMngr)
		})
	}
	for i := range el.exprs {
		ii := i
		execMngr.AddNextExec(el.exprs[ii], func() error {
			result.Add(results[ii].ObjectList[0])
			return nil
		})
	}
	return nil
}
