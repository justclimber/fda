package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewNumInt(id, value int64) *NumInt {
	return &NumInt{
		id:    id,
		value: value,
	}
}

type NumInt struct {
	id    int64
	value int64
}

func (n *NumInt) ID() int64            { return n.id }
func (n *NumInt) NodeKey() ast.NodeKey { return ast.KeyNumInt }

func (n *NumInt) Exec(_ *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(n, func() error {
		result.Add(&object.ObjInteger{Value: n.value})
		return nil
	})
	return nil
}
