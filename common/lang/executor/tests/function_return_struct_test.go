package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestFunctionReturnStruct(t *testing.T) {
	functionName := "testFunc"
	structName := "testStruct"
	varName1 := "a"
	testInt1 := int64(3)
	fieldName := "x"

	astStruct, structDefinition := getTestStructAstAndDefinition(t, testStruct{
		name: structName,
		fields: []testStructField{
			{
				name:      fieldName,
				fieldType: object.TypeInt,
				value:     ast.NewNumInt(0, testInt1),
			},
		},
	})
	definition := object.NewFunctionDefinition(
		functionName,
		"test",
		nil,
		[]*object.VarAndType{
			object.NewVarAndType(varName1, structDefinition.Type()),
		},
	)
	function := ast.NewFunction(
		0,
		definition,
		ast.NewStatementsBlock(0, []ast.Stmt{
			ast.NewVoidedExpression(
				0,
				ast.NewAssignment(0, []*ast.Identifier{ast.NewIdentifier(0, varName1)}, astStruct),
			),
		}),
	)
	functionCall := ast.NewFunctionCall(0, function, nil)

	packageAst := ast.NewPackage()
	packageAst.RegisterFunctionDefinition(definition)
	packagist := executor.NewPackagist(packageAst)
	env := environment.NewEnvironment()
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)

	res, err := ex.ExecAll(env, functionCall)
	require.NoError(t, err)
	require.NotEmpty(t, res.ObjectList)

	objStruct, ok := res.ObjectList[0].(*object.ObjStruct)
	require.True(t, ok, "check object is *object.ObjStruct")
	require.NotEmpty(t, objStruct.Fields, "check emptiness of struct fields")

	testObjectAsNumInt(t, objStruct.Fields[fieldName], testInt1)
}
