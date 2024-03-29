package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
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

func (el *NamedExpressionList) Exec(env *environment.Environment, result *object.NamedResult, execMngr execManager) error {
	results := make(map[string]*object.Result, len(el.exprs))
	for name, _ := range el.exprs {
		tName := name
		results[tName] = object.NewResult()
		execMngr.AddNextExec(el.exprs[tName], func() error {
			return el.exprs[tName].Exec(env, results[tName], execMngr)
		})
	}
	for name, _ := range el.exprs {
		tName := name
		execMngr.AddNextExec(el.exprs[tName], func() error {
			result.ObjectList[tName] = results[tName].ObjectList[0]
			return nil
		})
	}
	return nil
}
