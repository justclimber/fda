package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/errors"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
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

func (i *Identifier) Exec(env *environment.Environment, result *object.Result, _ execManager) error {
	if val, ok := env.Get(i.name); ok {
		result.Add(val)
		return nil
	}

	return errors.NewRuntimeError(i, errors.ErrorTypeIdentifierNotFound)
}
