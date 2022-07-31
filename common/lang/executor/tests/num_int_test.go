package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestNumInt_Exec(t *testing.T) {
	expectedInt := int64(4)
	numInt := ast.NewNumInt(expectedInt)

	env := environment.NewEnvironment()
	res := object.NewResult()
	packagist := executor.NewPackagist(nil)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := numInt.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	require.NoError(t, err, "check error from fn exec")
	testResultAsNumInt(t, res, expectedInt, 0)
}

func testResultAsNumInt(t *testing.T, res *object.Result, expectedInt int64, index int) {
	t.Helper()
	require.NotEmpty(t, res.ObjectList, "check result emptiness")

	testObjectAsNumInt(t, res.ObjectList[index], expectedInt)
}

func testObjectAsNumInt(t *testing.T, obj object.Object, expectedInt int64) {
	t.Helper()
	objInt, ok := obj.(*object.ObjInteger)
	require.True(t, ok, "check obj type")

	require.Equal(t, expectedInt, objInt.Value)
}
