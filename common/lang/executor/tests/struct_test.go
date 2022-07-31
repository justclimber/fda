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
	varName1, varName2 := "a", "b"
	fields := ast.NewAssignment(
		ast.NewIdentifierList([]string{varName1, varName2}),
		ast.NewExpressionList([]ast.Expr{
			ast.NewNumInt(expectedInt1),
			ast.NewNumInt(expectedInt2),
		}),
	)
	testStructName := "abc"
	astStruct := ast.NewStruct(testStructName, fields)

	structDefinitionFields := []*ast.VarAndType{
		ast.NewVarAndType(varName1, "int"),
		ast.NewVarAndType(varName2, "int"),
	}
	structDefinition := ast.NewStructDefinition(testStructName, structDefinitionFields)
	packageAst := ast.NewPackage(nil)
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
	testObjectAsNumInt(t, structObj.Fields[varName1], expectedInt1)
	testObjectAsNumInt(t, structObj.Fields[varName2], expectedInt2)
}
