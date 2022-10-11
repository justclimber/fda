package ast

import (
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

func (sb *StatementsBlock) Draw(r Renderer) {
	for _, statement := range sb.statements {
		statement.Draw(r)
		r.NewLine()
	}
}
