package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	execAst "github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
	"github.com/justclimber/fda/common/lang/validator"
	"github.com/justclimber/fda/common/lang/validator/ast"
)

func TestFunction(t *testing.T) {
	functionName := "testFunc"
	varName1, varName2 := "a", "b"
	inputVarName1, inputVarName2 := "inA", "inB"
	testInt1, testInt2, testInt3 := int64(3), int64(10), int64(2)
	definition := object.NewFunctionDefinition(
		functionName,
		"test",
		[]*object.VarAndType{
			object.NewVarAndType(inputVarName1, object.TypeInt),
			object.NewVarAndType(inputVarName2, object.TypeInt),
		},
		[]*object.VarAndType{
			object.NewVarAndType(varName1, object.TypeInt),
			object.NewVarAndType(varName2, object.TypeInt),
		},
	)
	function := ast.NewFunction(
		0,
		definition,
		ast.NewStatementsBlock(0, []ast.Stmt{
			ast.NewVoidedExpression(0, ast.NewAssignment(
				0,
				// a =
				[]*ast.Identifier{ast.NewIdentifier(0, varName1)},
				ast.NewNumInt(0, testInt1),
			)),
			ast.NewVoidedExpression(0, ast.NewAssignment(
				0,
				[]*ast.Identifier{ast.NewIdentifier(0, varName2), ast.NewIdentifier(0, "c")},
				ast.NewExpressionList(0, []ast.Expr{
					ast.NewNumInt(0, testInt2),
					ast.NewNumInt(0, 20),
				}),
			)),
		}),
	)
	functionCall := ast.NewFunctionCall(0, function, ast.NewNamedExpressionList(0, map[string]ast.Expr{
		inputVarName1: ast.NewNumInt(0, testInt3),
		inputVarName2: ast.NewNumInt(0, testInt3),
	}))

	envForValidation := validator.NewEnvironment()
	_, resAst, err := functionCall.Check(envForValidation, struct{}{})
	require.NoError(t, err, "check error after ast validation")

	functionCalForExec, ok := resAst.(*execAst.FunctionCall)
	require.True(t, ok, "check type is *FunctionCall")

	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(execQueue)

	envForExec := environment.NewEnvironment()
	res, err := ex.ExecAll(envForExec, functionCalForExec)
	require.NoError(t, err)
	require.NotEmpty(t, res.ObjectList)

	testObjectAsNumInt(t, res.ObjectList[0], testInt1)
	testObjectAsNumInt(t, res.ObjectList[1], testInt2)
}
