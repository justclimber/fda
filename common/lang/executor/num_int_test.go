package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/fdalang"
)

func TestNumInt_Exec(t *testing.T) {
	expectedInt := int64(4)
	ast := NewNumInt(expectedInt)

	env := fdalang.NewEnvironment()
	res := NewResult()
	execQueue := NewExecFnList()
	err := ast.Exec(env, res, execQueue)
	require.NoError(t, err, "check error from exec")

	err = execQueue.Current().fn()
	require.NoError(t, err, "check error from fn exec")
	testResultAsNumInt(t, res, expectedInt, 0)
}

func testResultAsNumInt(t *testing.T, res *Result, expectedInt int64, index int) {
	t.Helper()
	require.NotEmpty(t, res.objectList, "check result emptiness")

	testObjectAsNumInt(t, res.objectList[index], expectedInt)
}

func testObjectAsNumInt(t *testing.T, obj fdalang.Object, expectedInt int64) {
	t.Helper()
	objInt, ok := obj.(*fdalang.ObjInteger)
	require.True(t, ok, "check obj type")

	require.Equal(t, expectedInt, objInt.Value)
}
