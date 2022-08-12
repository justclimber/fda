package ast

import (
	"github.com/justclimber/fda/common/lang/ast"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
	"github.com/justclimber/fda/common/lang/validator/result"
)

type validationManager interface{}

type ValidatorEnv interface {
	Set(name string, objType object.Type)
	Check(name string, objType object.Type) bool
	Get(name string) (object.Type, bool)
	NewEnclosedEnvironment() ValidatorEnv
}

type Stmt interface {
	ast.Node
	Check(env ValidatorEnv, validMngr validationManager) (execAst.Stmt, error)
}

type Expr interface {
	ast.Node
	Check(env ValidatorEnv, execMngr validationManager) (*result.Result, execAst.Expr, error)
}
