package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestAssignment_Exec(t *testing.T) {
	expectedInt := int64(44)
	varName := "a"
	assignment := ast.NewAssignment(
		0,
		[]*ast.Identifier{ast.NewIdentifier(0, varName)},
		ast.NewNumInt(0, expectedInt),
	)

	env := environment.NewEnvironment()
	res := object.NewResult()
	packagist := executor.NewPackagist(nil)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := assignment.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	testResultAsNumInt(t, res, expectedInt, 0)

	obj, ok := env.Get(varName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, expectedInt)
}

func TestAssignment_Exec_Multiple(t *testing.T) {
	expectedInt1, expectedInt2 := int64(44), int64(55)
	varName1, varName2 := "a", "b"
	assignment := ast.NewAssignment(
		0,
		[]*ast.Identifier{ast.NewIdentifier(0, varName1), ast.NewIdentifier(0, varName2)},
		ast.NewExpressionList(0, []ast.Expr{
			ast.NewNumInt(0, expectedInt1),
			ast.NewNumInt(0, expectedInt2),
		}),
	)

	env := environment.NewEnvironment()
	res := object.NewResult()
	packagist := executor.NewPackagist(nil)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := assignment.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	testResultAsNumInt(t, res, expectedInt1, 0)
	testResultAsNumInt(t, res, expectedInt2, 1)

	obj1, ok := env.Get(varName1)
	require.True(t, ok, "check existence var1 in env")
	testObjectAsNumInt(t, obj1, expectedInt1)

	obj2, ok := env.Get(varName2)
	require.True(t, ok, "check existence var2 in env")
	testObjectAsNumInt(t, obj2, expectedInt2)
}
