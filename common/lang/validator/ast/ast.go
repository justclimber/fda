package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

type validationManager interface{}

type Stmt interface {
	ast.Node
	Exec(env *environment.Environment, validMngr validationManager) (execAst.Stmt, error)
}

type Expr interface {
	ast.Node
	Exec(env *environment.Environment, execMngr validationManager) (*object.Result, execAst.Expr, error)
}
