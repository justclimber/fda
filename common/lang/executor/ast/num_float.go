package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewNumFloat(value float64) *NumFloat {
	return &NumFloat{
		value: value,
	}
}

type NumFloat struct {
	id    int64
	value float64
}

func (n *NumFloat) ID() int64            { return n.id }
func (n *NumFloat) NodeKey() ast.NodeKey { return ast.KeyNumFloat }

func (n *NumFloat) Exec(_ *environment.Environment, result *object.Result, execMngr execManager) error {
	execMngr.AddNextExec(n, func() error {
		result.Add(&object.ObjFloat{Value: n.value})
		return nil
	})
	return nil
}
