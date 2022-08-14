package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
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
	errContainer := errors.NewErrContainer(a)
	value, exprAst, err := a.value.Check(env, validMngr)
	if err != nil {
		return nil, nil, errContainer.Wrap(err)
	}
	if value.Count() != len(a.left) {
		return nil, nil, errContainer.Wrap(errors.NewErrResultsCountMismatch(a, len(a.left), value.Count()))
	}
	identAst := make([]*execAst.Identifier, len(a.left))
	for i := range a.left {
		varName := a.left[i].name
		if !isExistedVarTypeMatched(a.left[i], env, varName, value.GetByIndex(i), errContainer) {
			continue
		}
		env.Set(varName, value.GetByIndex(i))
		identAst[i] = execAst.NewIdentifier(a.left[i].id, varName)
	}
	if errContainer.NotEmpty() {
		return nil, nil, errContainer
	}
	return value, execAst.NewAssignment(a.id, identAst, exprAst), nil
}

func isExistedVarTypeMatched(
	node ast.Node,
	env ValidatorEnv,
	varName string,
	expectedType object.Type,
	errContainer *errors.ErrContainer,
) bool {
	if objType, existed := env.GetRecursive(varName); existed && objType != expectedType {
		errContainer.Add(errors.NewErrTypesMismatchOnAssignment(node, varName, expectedType, objType))
		return false
	}
	return true
}
