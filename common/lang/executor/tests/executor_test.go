package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func TestExecutor(t *testing.T) {
	functionName := "main"
	function := ast.NewFunctionDefinition(
		functionName,
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
		nil,
		nil,
	)
	packageAst := ast.NewPackage()
	packageAst.RegisterFunctionDefinition(function)
	packagist := executor.NewPackagist(packageAst)
	env := environment.NewEnvironment()
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)

	functionCall := ast.NewFunctionCall(functionName, nil)
	_, err := ex.Exec(env, functionCall)
	require.NoError(t, err)
	env.Print()
}
