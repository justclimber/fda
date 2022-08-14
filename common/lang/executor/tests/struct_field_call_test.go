package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestStructFieldCall_Exec(t *testing.T) {
	structVarName := "s"
	testVarName := "q"
	structName := "abc"
	int1, int2 := int64(44), int64(55)
	fieldName1, fieldName2 := "a", "b"
	astStruct, _ := getTestStructAstAndDefinition(t, testStruct{
		name: structName,
		fields: []testStructField{
			{
				name:      fieldName1,
				fieldType: object.TypeInt,
				value:     ast.NewNumInt(0, int1),
			},
			{
				name:      fieldName2,
				fieldType: object.TypeInt,
				value:     ast.NewNumInt(0, int2),
			},
		},
	})

	astCode := ast.NewStatementsBlock(0, []ast.Stmt{
		ast.NewVoidedExpression(0, ast.NewAssignment(0, []*ast.Identifier{ast.NewIdentifier(0, structVarName)}, astStruct)),
		ast.NewVoidedExpression(0, ast.NewAssignment(0, []*ast.Identifier{ast.NewIdentifier(0, testVarName)}, ast.NewStructFieldCall(
			fieldName1,
			ast.NewIdentifier(0, structVarName),
		))),
	})
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(execQueue)
	env := environment.NewEnvironment()
	err := astCode.Exec(env, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)

	obj, ok := env.Get(testVarName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, int1)
}
