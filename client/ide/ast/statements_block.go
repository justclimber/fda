package ast

import (
	"fmt"

	"github.com/justclimber/fda/common/lang/ast"
)

func NewStatementsBlock(id int64, stmts []Stmt) *StatementsBlock {
	return &StatementsBlock{
		id:         id,
		statements: stmts,
	}
}

type StatementsBlock struct {
	id         int64
	statements []Stmt
}

func (sb *StatementsBlock) ID() int64            { return sb.id }
func (sb *StatementsBlock) NodeKey() ast.NodeKey { return ast.KeyStatementsBlock }

func (sb *StatementsBlock) Draw(r Renderer, slug string) {
	endNodeFunc := r.StartSiblingNode(sb, slug)
	r.StartContainerNode()
	for i, statement := range sb.statements {
		statement.Draw(r, fmt.Sprintf("%s stmt_%d", slug, i))

		// don't make a new line for last statement because of closing bracket and decreased indent
		if i != len(sb.statements)-1 {
			r.NewLine()
		}
	}
	r.EndContainerNode()
	endNodeFunc()
}
