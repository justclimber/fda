package ast

import (
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
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

func (i *Identifier) ID() int64        { return i.id }
func (i *Identifier) NodeKey() NodeKey { return KeyIdentifier }

func (i *Identifier) Exec(env *environment.Environment, result *object.Result, _ execManager) error {
	if val, ok := env.Get(i.value); ok {
		result.Add(val)
		return nil
	}

	return NewRuntimeError(i, ErrorTypeIdentifierNotFound)
}
