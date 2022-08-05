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
	inputVarName1, inputVarName2 := "inA", "inB"
	testInt1, testInt2, testInt3 := int64(3), int64(10), int64(2)
	function := ast.NewFunctionDefinition(
		functionName,
		ast.NewStatementsBlock([]ast.Stmt{
			ast.NewVoidedExpression(
				ast.NewAssignment(
					// a =
					ast.NewIdentifierList([]string{varName1}),
					// inA + 3
					ast.NewArithmeticOperation(
						ast.NewIdentifier(inputVarName1),
						ast.NewNumInt(testInt1),
						object.OperatorAddition,
					),
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
		[]*ast.VarAndType{
			ast.NewVarAndType(inputVarName1, object.TypeInt),
			ast.NewVarAndType(inputVarName2, object.TypeInt),
		},
		[]*ast.VarAndType{
			ast.NewVarAndType(varName1, object.TypeInt),
			ast.NewVarAndType(varName2, object.TypeInt),
		},
	)
	functionCall := ast.NewFunctionCall(
		functionName,
		ast.NewNamedExpressionList(map[string]ast.Expr{
			inputVarName1: ast.NewNumInt(testInt3), // inA = 2
			inputVarName2: ast.NewNumInt(testInt3),
		}),
	)

	packageAst := ast.NewPackage()
	packageAst.RegisterFunctionDefinition(function)
	packagist := executor.NewPackagist(packageAst)
	env := environment.NewEnvironment()
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)

	res, err := ex.Exec(env, functionCall)
	require.NoError(t, err)
	require.NotEmpty(t, res.ObjectList)

	testObjectAsNumInt(t, res.ObjectList[0], testInt1+testInt3)
	testObjectAsNumInt(t, res.ObjectList[1], testInt2)
}
