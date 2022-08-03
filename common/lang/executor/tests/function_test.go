package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestFunction(t *testing.T) {
	functionName := "testFunc"
	varName1, varName2 := "a", "b"
	testInt1, testInt2 := int64(3), int64(10)
	function := ast.NewFunctionDefinition(
		functionName,
		ast.NewStatementsBlock([]ast.Stmt{
			ast.NewVoidedExpression(
				ast.NewAssignment(
					ast.NewIdentifierList([]string{varName1}),
					ast.NewNumInt(testInt1),
				),
			),
			ast.NewVoidedExpression(
				ast.NewAssignment(
					ast.NewIdentifierList([]string{varName2, "c"}),
					ast.NewExpressionList([]ast.Expr{
						ast.NewNumInt(testInt2),
						ast.NewNumInt(20),
					}),
				),
			),
		}),
		nil,
		[]*ast.VarAndType{
			ast.NewVarAndType(varName1, object.TypeInt),
			ast.NewVarAndType(varName2, object.TypeInt),
		},
	)
	packageAst := ast.NewPackage()
	packageAst.RegisterFunctionDefinition(function)
	packagist := executor.NewPackagist(packageAst)
	env := environment.NewEnvironment()
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)

	functionCall := ast.NewFunctionCall(functionName)
	res, err := ex.Exec(env, functionCall)
	require.NoError(t, err)
	require.NotEmpty(t, res.ObjectList)

	testObjectAsNumInt(t, res.ObjectList[0], testInt1)
	testObjectAsNumInt(t, res.ObjectList[1], testInt2)
}
