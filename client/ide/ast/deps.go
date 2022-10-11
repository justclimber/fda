package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

type DrawableNode interface {
	ast.Node
	Draw(r Renderer)
}

type Stmt interface {
	DrawableNode
}

type Expr interface {
	DrawableNode
}

type TextType int

const (
	TypeSystemSymbols = TextType(iota + 1)
	TypeIdentifier
	TypeNumbers
)

type Renderer interface {
	DrawAssignment()
	DrawArgDelimiter()
	NewLine()
	IndentIncrease()
	IndentDecrease()
	DrawText(name string, t TextType)
}
