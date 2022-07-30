package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/fdalang"
)

func TestExecutor(t *testing.T) {
	packageAst := NewPackage(
		NewFunction(
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
		),
	)
	env := fdalang.NewEnvironment()
	ex := NewExecutor(env, packageAst)
	err := ex.Exec()
	require.NoError(t, err)
	ex.DebugPrint()
}
