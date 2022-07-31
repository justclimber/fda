package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func testNext(t *testing.T, execQueue *executor.ExecFnList, times int) {
	t.Helper()
	for i := 0; i < times; i++ {
		_, err := execQueue.ExecNext()
		require.NoError(t, err, "check error from fn exec")
	}
}

func testNextAll(t *testing.T, execQueue *executor.ExecFnList) {
	t.Helper()
	var err error
	hasNext := true
	for hasNext {
		hasNext, err = execQueue.ExecNext()
		require.NoError(t, err, "check error from fn exec")
	}
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

	assert.Equal(t, expectedInt, objInt.Value)
}

func getTestStruct(t *testing.T, testStructName string) (*ast.Struct, *ast.StructDefinition) {
	t.Helper()
	expectedInt1, expectedInt2 := int64(44), int64(55)
	varName1, varName2 := "a", "b"
	fields := ast.NewNamedExpressionList(map[string]ast.Expr{
		varName1: ast.NewNumInt(expectedInt1),
		varName2: ast.NewNumInt(expectedInt2),
	})
	astStruct := ast.NewStruct(testStructName, fields)

	structDefinitionFields := map[string]*ast.VarAndType{
		varName1: ast.NewVarAndType(varName1, "int"),
		varName2: ast.NewVarAndType(varName2, "int"),
	}
	structDefinition := ast.NewStructDefinition(testStructName, structDefinitionFields)

	return astStruct, structDefinition
}
