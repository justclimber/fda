package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type execManager interface {
	AddNextExec(node ast.Node, fn func() error)
	MainPackage() *Package
}

type Stmt interface {
	ast.Node
	Exec(env *environment.Environment, execMngr execManager) error
}

type Expr interface {
	ast.Node
	Exec(env *environment.Environment, result *object.Result, execMngr execManager) error
}
