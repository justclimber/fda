package executor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStruct_Exec(t *testing.T) {
	expectedInt1, expectedInt2 := int64(44), int64(55)
	varName1, varName2 := "a", "b"
	fields := NewAssignment(
		NewIdentifierList([]string{varName1, varName2}),
		NewExpressionList([]Expr{
			NewNumInt(expectedInt1),
			NewNumInt(expectedInt2),
		}),
	)
	testStructName := "abc"
	ast := NewStruct(testStructName, fields)

	structDefinitionFields := []*VarAndType{
		{
			varType: "int",
			varName: varName1,
		},
		{
			varType: "int",
			varName: varName2,
		},
	}
	structDefinition := NewStructDefinition(testStructName, structDefinitionFields)
	packageAst := NewPackage(nil)
	packageAst.RegisterStructDefinition(structDefinition)
	packagist := NewPackagist(packageAst)
	execQueue := NewExecFnList()
	res := NewResult()
	ex := NewExecutor(packagist, execQueue)
	env := NewEnvironment()
	err := ast.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	require.NotEmpty(t, res.objectList, "check result emptiness")

	structObj, ok := res.objectList[0].(*ObjStruct)
	require.True(t, ok, "check obj type")

	require.NotEmpty(t, structObj.Fields, "check struct fields emptiness")
	testObjectAsNumInt(t, structObj.Fields[varName1], expectedInt1)
	testObjectAsNumInt(t, structObj.Fields[varName2], expectedInt2)
}
