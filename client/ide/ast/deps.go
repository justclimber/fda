package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
)

type Stmt interface {
	ast.Node
}

type Expr interface {
	ast.Node
}
