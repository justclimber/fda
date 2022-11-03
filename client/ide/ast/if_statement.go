package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

func NewIfStatement(id int64, condition Expr, trueBranch *StatementsBlock, falseBranch *StatementsBlock) *IfStatement {
	return &IfStatement{
		id:          id,
		condition:   condition,
		trueBranch:  trueBranch,
		falseBranch: falseBranch,
	}
}

type IfStatement struct {
	id          int64
	condition   Expr
	trueBranch  *StatementsBlock
	falseBranch *StatementsBlock
}

func (is *IfStatement) ID() int64            { return is.id }
func (is *IfStatement) NodeKey() ast.NodeKey { return ast.KeyIfStatement }

func (is *IfStatement) Draw(r Renderer, _ string) {
	endNodeFunc := r.StartSiblingNode(is, "if")
	r.StartContainerNode()
	r.DrawIfStart()
	is.condition.Draw(r, "condition")
	r.DrawIfMid()
	is.trueBranch.Draw(r, "true branch")
	if is.falseBranch != nil {
		r.DrawIfElse()
		is.falseBranch.Draw(r, "else branch")
	}
	r.DrawIfEnd()
	endNodeFunc()
	r.EndContainerNode()
}
