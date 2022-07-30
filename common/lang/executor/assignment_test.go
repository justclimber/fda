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
	testResultAsNumInt(t, res, expectedInt, 0)
	testNext(t, execQueue, 1)

	obj, ok := env.Get(varName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, expectedInt)
}

func TestAssignment_Exec_Multiple(t *testing.T) {
	expectedInt1, expectedInt2 := int64(44), int64(55)
	varName1, varName2 := "a", "b"
	ast := NewAssignment(
		NewIdentifierList([]string{varName1, varName2}),
		NewExpressionList([]Expr{
			NewNumInt(expectedInt1),
			NewNumInt(expectedInt2),
		}),
	)

	env := fdalang.NewEnvironment()
	res := NewResult()
	execQueue := NewExecFnList()
	err := ast.Exec(env, res, execQueue)
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

func testNext(t *testing.T, execQueue *ExecFnList, times int) {
	t.Helper()
	for i := 0; i < times; i++ {
		_, err := execQueue.Next()
		require.NoError(t, err, "check error from fn exec")
	}
}

func testNextAll(t *testing.T, execQueue *ExecFnList) {
	t.Helper()
	var err error
	hasNext := true
	for hasNext {
		hasNext, err = execQueue.Next()
		require.NoError(t, err, "check error from fn exec")
	}
}
