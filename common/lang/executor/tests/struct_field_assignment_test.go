package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestStructFieldAssignment(t *testing.T) {
	structVarName := "s"
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
	testInt := int64(123)

	astCode := ast.NewStatementsBlock([]ast.Stmt{
		ast.NewVoidedExpression(
			ast.NewAssignment(
				ast.NewIdentifierList([]string{structVarName}),
				astStruct,
			),
		),
		ast.NewVoidedExpression(
			ast.NewStructFieldAssignment(
				[]*ast.StructFieldIdentifier{
					ast.NewStructFieldIdentifier(
						fieldName1,
						ast.NewIdentifier(structVarName),
					),
				},
				ast.NewNumInt(testInt),
			),
		),
	})
	packageAst := ast.NewPackage(nil)
	packageAst.RegisterStructDefinition(structDefinition)
	packagist := executor.NewPackagist(packageAst)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	env := environment.NewEnvironment()
	err := astCode.Exec(env, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)

	obj, ok := env.Get(structVarName)
	require.True(t, ok, "check existence var in env")

	structObj, ok := obj.(*object.ObjStruct)
	require.True(t, ok, "check obj type")

	require.NotEmpty(t, structObj.Fields, "check struct fields emptiness")
	testObjectAsNumInt(t, structObj.Fields[fieldName1], testInt)
}
