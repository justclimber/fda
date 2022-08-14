package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/validator/result"
)

func NewIdentifier(id int64, name string) *Identifier {
	return &Identifier{
		id:   id,
		name: name,
	}
}

type Identifier struct {
	id   int64
	name string
}

func (i *Identifier) ID() int64            { return i.id }
func (i *Identifier) NodeKey() ast.NodeKey { return ast.KeyIdentifier }

func (i *Identifier) Check(env ValidatorEnv, _ validationManager) (*result.Result, execAst.Expr, error) {
	if objType, ok := env.GetRecursive(i.name); ok {
		return result.NewSingleResult(objType), execAst.NewIdentifier(i.id, i.name), nil
	}

	return nil, nil, errors.NewErrIdentifierNotFound(i, i.name)
}
