package ast

import (
	"strconv"

	"github.com/justclimber/fda/common/lang/ast"
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

func (n *NumInt) Draw(r Renderer) {
	r.DrawText(strconv.FormatInt(n.value, 10), TypeNumbers)
}
