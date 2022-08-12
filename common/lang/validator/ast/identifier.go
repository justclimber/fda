package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/validator/result"
)

func NewIdentifier(id int64, value string) *Identifier {
	return &Identifier{
		id:    id,
		value: value,
	}
}

type Identifier struct {
	id    int64
	value string
}

func (i *Identifier) ID() int64            { return i.id }
func (i *Identifier) NodeKey() ast.NodeKey { return ast.KeyIdentifier }

func (i *Identifier) Check(
	env ValidatorEnv,
	_ validationManager,
) (*result.Result, execAst.Expr, error) {
	if objType, ok := env.Get(i.value); ok {
		res := result.NewResult()
		res.Add(objType)
		return res, execAst.NewIdentifier(i.id, i.value), nil
	}

	return nil, nil, errors.NewValidationError(i, errors.ErrorTypeIdentifierNotFound)
}
