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

type testStruct struct {
	name   string
	fields []testStructField
}

type testStructField struct {
	name      string
	fieldType object.Type
	value     ast.Expr
}

func getTestStructAstAndDefinition(t *testing.T, testStruct testStruct) (*ast.Struct, *object.StructDefinition) {
	t.Helper()

	f := map[string]ast.Expr{}
	for _, field := range testStruct.fields {
		f[field.name] = field.value
	}

	structDefinitionFields := map[string]*object.VarAndType{}
	for _, field := range testStruct.fields {
		structDefinitionFields[field.name] = object.NewVarAndType(field.name, field.fieldType)
	}
	structDefinition := object.NewStructDefinition(testStruct.name, "test", structDefinitionFields)

	astStruct := ast.NewStruct(0, structDefinition, ast.NewNamedExpressionList(0, f))

	return astStruct, structDefinition
}
