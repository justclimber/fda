package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
	"github.com/justclimber/fda/common/lang/validator/result"
)

func NewNumInt(id int64, value int64) *NumInt {
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

func (n *NumInt) Check(_ ValidatorEnv, _ validationManager) (*result.Result, execAst.Expr, error) {
	return result.NewSingleResult(object.TypeInt), execAst.NewNumInt(n.id, n.value), nil
}
