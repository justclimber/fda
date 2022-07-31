package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func TestExecutor(t *testing.T) {
	function := ast.NewFunction(
		ast.NewStatementsBlock([]ast.Stmt{
			ast.NewVoidedExpression(
				ast.NewAssignment(
					ast.NewIdentifierList([]string{"a"}),
					ast.NewUnaryMinus(
						ast.NewNumInt(3),
					),
				),
			),
			ast.NewVoidedExpression(
				ast.NewAssignment(
					ast.NewIdentifierList([]string{"b", "c"}),
					ast.NewExpressionList([]ast.Expr{
						ast.NewNumInt(10),
						ast.NewNumInt(20),
					}),
				),
			),
		}),
	)
	packageAst := ast.NewPackage(function)
	packagist := executor.NewPackagist(packageAst)
	env := environment.NewEnvironment()
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := ex.Exec(env, function)
	require.NoError(t, err)
	env.Print()
}
