package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewBool(value bool) *Bool {
	return &Bool{
		value: value,
	}
}

type Bool struct {
	id    int64
	value bool
}

func (b *Bool) ID() int64            { return b.id }
func (b *Bool) NodeKey() ast.NodeKey { return ast.KeyBool }

func (b *Bool) Exec(_ *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(b, func() error {
		result.Add(object.ToReservedBoolObj(b.value))
		return nil
	})
	return nil
}
