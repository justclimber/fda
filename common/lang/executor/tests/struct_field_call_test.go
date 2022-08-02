package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func TestStructFieldCall_Exec(t *testing.T) {
	structVarName := "s"
	testVarName := "q"
	structName := "abc"
	int1, int2 := int64(44), int64(55)
	fieldName1, fieldName2 := "a", "b"
	astStruct, structDefinition := getTestStruct(t, testStruct{
		name: structName,
		fields: []testStructField{
			{
				name:      fieldName1,
				fieldType: "int",
				intValue:  int1,
			},
			{
				name:      fieldName2,
				fieldType: "int",
				intValue:  int2,
			},
		},
	})

	astCode := ast.NewStatementsBlock([]ast.Stmt{
		ast.NewVoidedExpression(
			ast.NewAssignment(
				ast.NewIdentifierList([]string{structVarName}),
				astStruct,
			),
		),
		ast.NewVoidedExpression(
			ast.NewAssignment(
				ast.NewIdentifierList([]string{testVarName}),
				ast.NewStructFieldCall(
					fieldName1,
					ast.NewIdentifier(structVarName),
				),
			),
		),
	})
	packageAst := ast.NewPackage()
	packageAst.RegisterStructDefinition(structDefinition)
	packagist := executor.NewPackagist(packageAst)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	env := environment.NewEnvironment()
	err := astCode.Exec(env, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)

	obj, ok := env.Get(testVarName)
	require.True(t, ok, "check existence var in env")
	testObjectAsNumInt(t, obj, int1)
}
