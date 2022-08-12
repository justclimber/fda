package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/validator/result"
)

func NewAssignment(id int64, left []*Identifier, value Expr) *Assignment {
	return &Assignment{
		id:    id,
		left:  left,
		value: value,
	}
}

type Assignment struct {
	id    int64
	left  []*Identifier
	value Expr
}

func (a *Assignment) ID() int64            { return a.id }
func (a *Assignment) NodeKey() ast.NodeKey { return ast.KeyAssignment }

func (a *Assignment) Check(env ValidatorEnv, validMngr validationManager) (*result.Result, execAst.Expr, error) {
	value, exprAst, err := a.value.Check(env, validMngr)
	if err != nil {
		return nil, nil, err
	}
	identAst := make([]*execAst.Identifier, len(a.left))
	for i := range a.left {
		varName := a.left[i].value
		env.Set(varName, value.GetByIndex(i))
		identAst[i] = execAst.NewIdentifier(a.left[i].id, varName)
	}
	return value, execAst.NewAssignment(a.id, identAst, exprAst), nil
}
