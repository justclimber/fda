package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewAssignment(id int64, left []*Identifier, value Expr) *Assignment {
	return &Assignment{
		id:    id,
		left:  left,
		value: value,
	}
}

type Assignment struct {
	id    int64
	left  []*Identifier
	value Expr
}

func (a *Assignment) ID() int64            { return a.id }
func (a *Assignment) NodeKey() ast.NodeKey { return ast.KeyAssignment }

func (a *Assignment) Draw(r Renderer) {
	count := len(a.left)
	for i, identifier := range a.left {
		identifier.Draw(r)
		if i != count-1 {
			r.DrawArgDelimiter()
		}
	}
	r.DrawAssignment()
	a.value.Draw(r)
}
