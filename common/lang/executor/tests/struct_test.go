package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestStruct_Exec(t *testing.T) {
	expectedInt1, expectedInt2 := int64(44), int64(55)
	fieldName1, fieldName2 := "a", "b"
	testStructName := "abc"
	astStruct, structDefinition := getTestStructAstAndDefinition(t, testStruct{
		name: testStructName,
		fields: []testStructField{
			{
				name:      fieldName1,
				fieldType: object.TypeInt,
				value:     ast.NewNumInt(0, expectedInt1),
			},
			{
				name:      fieldName2,
				fieldType: object.TypeInt,
				value:     ast.NewNumInt(0, expectedInt2),
			},
		},
	})

	packageAst := ast.NewPackage()
	packageAst.RegisterStructDefinition(structDefinition)
	packagist := executor.NewPackagist(packageAst)
	execQueue := executor.NewExecFnList()
	res := object.NewResult()
	ex := executor.NewExecutor(packagist, execQueue)
	env := environment.NewEnvironment()
	err := astStruct.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	require.NotEmpty(t, res.ObjectList, "check result emptiness")

	structObj, ok := res.ObjectList[0].(*object.ObjStruct)
	require.True(t, ok, "check obj type")

	require.NotEmpty(t, structObj.Fields, "check struct fields emptiness")
	testObjectAsNumInt(t, structObj.Fields[fieldName1], expectedInt1)
	testObjectAsNumInt(t, structObj.Fields[fieldName2], expectedInt2)
}
