package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumInt_Exec(t *testing.T) {
	expectedInt := int64(4)
	ast := NewNumInt(expectedInt)

	env := NewEnvironment()
	res := NewResult()
	packagist := NewPackagist(nil)
	execQueue := NewExecFnList()
	ex := NewExecutor(packagist, execQueue)
	err := ast.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	require.NoError(t, err, "check error from fn exec")
	testResultAsNumInt(t, res, expectedInt, 0)
}

func testResultAsNumInt(t *testing.T, res *Result, expectedInt int64, index int) {
	t.Helper()
	require.NotEmpty(t, res.objectList, "check result emptiness")

	testObjectAsNumInt(t, res.objectList[index], expectedInt)
}

func testObjectAsNumInt(t *testing.T, obj Object, expectedInt int64) {
	t.Helper()
	objInt, ok := obj.(*ObjInteger)
	require.True(t, ok, "check obj type")

	require.Equal(t, expectedInt, objInt.Value)
}
