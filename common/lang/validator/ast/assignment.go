package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewAssignment(left []*Identifier, value Expr) *Assignment {
	return &Assignment{
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

func (a *Assignment) Exec(env *environment.Environment, validMngr validationManager) (*object.Result, execAst.Expr, error) {
	value, exprAst, err := a.value.Exec(env, validMngr)
	if err != nil {
		return nil, nil, err
	}
	identAst := make([]*execAst.Identifier, len(a.left))
	for i := range a.left {
		varName := a.left[i].value
		env.Set(varName, value.ObjectList[i])
		identAst[i] = execAst.NewIdentifier(a.left[i].id, varName)
	}
	return value, execAst.NewAssignment(a.id, identAst, exprAst), nil
}
