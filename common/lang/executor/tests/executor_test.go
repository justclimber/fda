package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestExecutor(t *testing.T) {
	functionName := "main"
	definition := object.NewFunctionDefinition(functionName, "test", nil, nil)
	function := ast.NewFunction(
		0,
		definition,
		ast.NewStatementsBlock(0, []ast.Stmt{
			ast.NewVoidedExpression(
				0,
				ast.NewAssignment(0, []*ast.Identifier{ast.NewIdentifier(0, "a")}, ast.NewUnaryMinus(
					ast.NewNumInt(3, 0),
				)),
			),
			ast.NewVoidedExpression(
				0,
				ast.NewAssignment(
					0,
					[]*ast.Identifier{ast.NewIdentifier(0, "b"), ast.NewIdentifier(0, "c")},
					ast.NewExpressionList(0, []ast.Expr{
						ast.NewNumInt(10, 0),
						ast.NewNumInt(20, 0),
					}),
				),
			),
		}),
	)
	packageAst := ast.NewPackage()
	packageAst.RegisterFunctionDefinition(definition)
	packagist := executor.NewPackagist(packageAst)
	env := environment.NewEnvironment()
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)

	functionCall := ast.NewFunctionCall(0, function, nil)
	_, err := ex.ExecAll(env, functionCall)
	require.NoError(t, err)
	env.Print()
}
