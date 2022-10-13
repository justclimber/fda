package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewIfStatement(condition Expr, trueBranch, falseBranch *StatementsBlock) *IfStatement {
	return &IfStatement{
		condition:   condition,
		trueBranch:  trueBranch,
		falseBranch: falseBranch,
	}
}

type IfStatement struct {
	id          int64
	condition   Expr
	trueBranch  *StatementsBlock
	falseBranch *StatementsBlock
}

func (is *IfStatement) ID() int64            { return is.id }
func (is *IfStatement) NodeKey() ast.NodeKey { return ast.KeyIfStatement }

func (is *IfStatement) Check(env ValidatorEnv, validMngr validationManager) (execAst.Stmt, error) {
	errContainer := errors.NewErrContainer(is)
	value, conditionAst, err := is.condition.Check(env, validMngr)
	if err != nil {
		errContainer.Add(err)
	} else if resultType := value.Get(); resultType != object.TypeBool {
		errContainer.Add(errors.NewErrIfExpressionResultMustBeBool(is.condition, resultType))
	}

	trueBranchAst, err := is.trueBranch.Check(env, validMngr)
	if err != nil {
		errContainer.Add(err)
	}

	var falseBranchAst *execAst.StatementsBlock
	if is.falseBranch != nil {
		falseBranchAst, err = is.falseBranch.Check(env, validMngr)
		if err != nil {
			errContainer.Add(err)
		}
	}
	if errContainer.NotEmpty() {
		return nil, errContainer
	}
	return execAst.NewIfStatement(conditionAst, trueBranchAst, falseBranchAst), nil
}
