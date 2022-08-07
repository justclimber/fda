package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func NewNumInt(id int64, value int64) *NumInt {
	return &NumInt{
		id:    value,
		value: value,
	}
}

type NumInt struct {
	id    int64
	value int64
}

func (n *NumInt) ID() int64            { return n.id }
func (n *NumInt) NodeKey() ast.NodeKey { return ast.KeyNumInt }

func (n *NumInt) Exec(_ *environment.Environment, _ validationManager) (*object.Result, execAst.Expr, error) {
	r := object.NewResult()
	r.Add(&object.ObjInteger{Value: n.value})
	return r, execAst.NewNumInt(n.id, n.value), nil
}
