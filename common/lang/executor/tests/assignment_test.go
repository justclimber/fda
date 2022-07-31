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
		ast.NewIdentifierList([]string{varName}),
		ast.NewNumInt(expectedInt),
	)

	env := environment.NewEnvironment()
	res := object.NewResult()
	packagist := executor.NewPackagist(nil)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := assignment.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNext(t, execQueue, 2)
	testResultAsNumInt(t, res, expectedInt, 0)
	testNext(t, execQueue, 1)

	obj, ok := env.Get(varName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, expectedInt)
}

func TestAssignment_Exec_Multiple(t *testing.T) {
	expectedInt1, expectedInt2 := int64(44), int64(55)
	varName1, varName2 := "a", "b"
	assignment := ast.NewAssignment(
		ast.NewIdentifierList([]string{varName1, varName2}),
		ast.NewExpressionList([]ast.Expr{
			ast.NewNumInt(expectedInt1),
			ast.NewNumInt(expectedInt2),
		}),
	)

	env := environment.NewEnvironment()
	res := object.NewResult()
	packagist := executor.NewPackagist(nil)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := assignment.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNext(t, execQueue, 7)
	testResultAsNumInt(t, res, expectedInt1, 0)
	testResultAsNumInt(t, res, expectedInt2, 1)
	testNext(t, execQueue, 2)

	obj1, ok := env.Get(varName1)
	require.True(t, ok, "check existence var1 in env")
	testObjectAsNumInt(t, obj1, expectedInt1)

	obj2, ok := env.Get(varName1)
	require.True(t, ok, "check existence var2 in env")
	testObjectAsNumInt(t, obj2, expectedInt1)
}
