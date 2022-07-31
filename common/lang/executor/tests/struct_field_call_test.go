package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func TestStructFieldCall_Exec(t *testing.T) {
	structVarName := "s"
	testVarName := "q"
	structName := "abc"
	astStruct, structDefinition := getTestStruct(t, structName)

	astCode := ast.NewStatementsBlock([]ast.Stmt{
		ast.NewVoidedExpression(
			ast.NewAssignment(
				ast.NewIdentifierList([]string{structVarName}),
				astStruct,
			),
		),
		ast.NewVoidedExpression(
			ast.NewAssignment(
				ast.NewIdentifierList([]string{testVarName}),
				ast.NewStructFieldCall(
					"a",
					ast.NewIdentifier(structVarName),
				),
			),
		),
	})
	packageAst := ast.NewPackage(nil)
	packageAst.RegisterStructDefinition(structDefinition)
	packagist := executor.NewPackagist(packageAst)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	env := environment.NewEnvironment()
	err := astCode.Exec(env, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)

	obj, ok := env.Get(testVarName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, 44)
}
