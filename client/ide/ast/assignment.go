package ast

import (
	"fmt"

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

func (a *Assignment) Draw(r Renderer, slug string) {
	endFunc := r.StartSiblingNode(a, slug+" assignment")
	r.StartContainerNode()
	count := len(a.left)
	for i, identifier := range a.left {
		identifier.Draw(r, fmt.Sprintf("assignment_ident_%d", i))
		if i != count-1 {
			r.DrawArgDelimiter()
		}
	}
	r.DrawAssignment()
	a.value.Draw(r, "assignment value")
	endFunc()
	r.EndContainerNode()
}
