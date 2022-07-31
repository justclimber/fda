package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecutor(t *testing.T) {
	function := NewFunction(
		NewStatementsBlock([]Stmt{
			NewVoidedExpression(
				NewAssignment(
					NewIdentifierList([]string{"a"}),
					NewUnaryMinus(
						NewNumInt(3),
					),
				),
			),
			NewVoidedExpression(
				NewAssignment(
					NewIdentifierList([]string{"b", "c"}),
					NewExpressionList([]Expr{
						NewNumInt(10),
						NewNumInt(20),
					}),
				),
			),
		}),
	)
	packageAst := NewPackage(function)
	packagist := NewPackagist(packageAst)
	env := NewEnvironment()
	execQueue := NewExecFnList()
	ex := NewExecutor(packagist, execQueue)
	err := ex.Exec(env, function)
	require.NoError(t, err)
	env.Print()
}
