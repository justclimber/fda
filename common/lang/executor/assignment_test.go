package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/fdalang"
)

func TestAssignment_Exec(t *testing.T) {
	expectedInt := int64(44)
	varName := "a"
	ast := NewAssignment(
		NewIdentifierList([]string{varName}),
		NewNumInt(expectedInt),
	)

	env := fdalang.NewEnvironment()
	res := NewResult()
	execQueue := NewExecFnList()
	err := ast.Exec(env, res, execQueue)
	require.NoError(t, err, "check error from exec")

	testNext(t, execQueue, 2)
	testResultAsNumInt(t, res, expectedInt)
	testNext(t, execQueue, 1)

	obj, ok := env.Get(varName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, expectedInt)
}

func testNext(t *testing.T, execQueue *ExecFnList, times int) {
	t.Helper()
	for i := 0; i < times; i++ {
		_, err := execQueue.Next()
		require.NoError(t, err, "check error from fn exec")
	}
}
