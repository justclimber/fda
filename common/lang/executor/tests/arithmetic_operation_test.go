package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestArithmeticOperation(t *testing.T) {
	tests := []struct {
		name             string
		left             ast.Expr
		right            ast.Expr
		operator         object.ArithmeticOperator
		expectedResultFn func(r *object.Result)
	}{
		{
			name:             "int_addition_1",
			left:             ast.NewNumInt(0, 3),
			right:            ast.NewNumInt(0, 7),
			operator:         object.OperatorAddition,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: 10}) },
		},
		{
			name:             "int_addition_with-minus-arg",
			left:             ast.NewNumInt(0, 3),
			right:            ast.NewNumInt(0, -7),
			operator:         object.OperatorAddition,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: -4}) },
		},
		{
			name:             "int_addition_with-zero",
			left:             ast.NewNumInt(0, 3),
			right:            ast.NewNumInt(0, 0),
			operator:         object.OperatorAddition,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: 3}) },
		},
		{
			name:             "int_subtraction_1",
			left:             ast.NewNumInt(0, 3),
			right:            ast.NewNumInt(0, 7),
			operator:         object.OperatorSubtraction,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: -4}) },
		},
		{
			name:             "int_subtraction_with-minus-arg",
			left:             ast.NewNumInt(0, 3),
			right:            ast.NewNumInt(0, -7),
			operator:         object.OperatorSubtraction,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: 10}) },
		},
		{
			name:             "int_subtraction_with-zero",
			left:             ast.NewNumInt(0, 0),
			right:            ast.NewNumInt(0, 3),
			operator:         object.OperatorSubtraction,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: -3}) },
		},
		{
			name:             "float_addition_1",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorAddition,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: 10}) },
		},
		{
			name:             "float_addition_with-minus-arg",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(-7),
			operator:         object.OperatorAddition,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: -4}) },
		},
		{
			name:             "float_addition_with-zero",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(0),
			operator:         object.OperatorAddition,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: 3}) },
		},
		{
			name:             "float_subtraction_1",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorSubtraction,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: -4}) },
		},
		{
			name:             "float_subtraction_with-minus-arg",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(-7),
			operator:         object.OperatorSubtraction,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: 10}) },
		},
		{
			name:             "float_subtraction_with-zero",
			left:             ast.NewNumFloat(0),
			right:            ast.NewNumFloat(3),
			operator:         object.OperatorSubtraction,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: -3}) },
		},
		{
			name:             "int_multiplication",
			left:             ast.NewNumInt(0, 2),
			right:            ast.NewNumInt(0, 3),
			operator:         object.OperatorMultiplication,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: 6}) },
		},
		{
			name:             "float_multiplication",
			left:             ast.NewNumFloat(2),
			right:            ast.NewNumFloat(3),
			operator:         object.OperatorMultiplication,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: 6}) },
		},
		{
			name:             "int_division",
			left:             ast.NewNumInt(0, 6),
			right:            ast.NewNumInt(0, 3),
			operator:         object.OperatorDivision,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjInteger{Value: 2}) },
		},
		{
			name:             "float_division",
			left:             ast.NewNumFloat(5),
			right:            ast.NewNumFloat(2),
			operator:         object.OperatorDivision,
			expectedResultFn: func(r *object.Result) { r.Add(&object.ObjFloat{Value: 2.5}) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arithmeticOperation := ast.NewArithmeticOperation(tt.left, tt.right, tt.operator)
			env := environment.NewEnvironment()

			expectedRes := object.NewResult()
			tt.expectedResultFn(expectedRes)

			actualRes := object.NewResult()
			execQueue := executor.NewExecFnList()
			ex := executor.NewExecutor(execQueue)
			err := arithmeticOperation.Exec(env, actualRes, ex)
			require.NoError(t, err, "check error from exec")

			testNextAll(t, execQueue)
			require.NoError(t, err, "check error from fn exec")
			require.Equal(t, expectedRes, actualRes, "check result is equal")
		})
	}
}
