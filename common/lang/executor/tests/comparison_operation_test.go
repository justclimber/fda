package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestComparisonOperation(t *testing.T) {
	tests := []struct {
		name             string
		left             ast.Expr
		right            ast.Expr
		operator         object.ComparisonOperator
		expectedResultFn func(r *object.Result)
	}{
		// equal
		{
			name:             "int_equal_false",
			left:             ast.NewNumInt(3),
			right:            ast.NewNumInt(7),
			operator:         object.OperatorEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "int_equal_true",
			left:             ast.NewNumInt(7),
			right:            ast.NewNumInt(7),
			operator:         object.OperatorEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		{
			name:             "float_equal_false",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "float_equal_true",
			left:             ast.NewNumFloat(7),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		{
			name:             "bool_equal_false",
			left:             ast.NewBool(true),
			right:            ast.NewBool(false),
			operator:         object.OperatorEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "bool_equal_true",
			left:             ast.NewBool(true),
			right:            ast.NewBool(true),
			operator:         object.OperatorEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		// not equal
		{
			name:             "int_not_equal_true",
			left:             ast.NewNumInt(3),
			right:            ast.NewNumInt(7),
			operator:         object.OperatorNotEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		{
			name:             "int_not_equal_false",
			left:             ast.NewNumInt(7),
			right:            ast.NewNumInt(7),
			operator:         object.OperatorNotEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "float_not_equal_true",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorNotEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		{
			name:             "float_not_equal_false",
			left:             ast.NewNumFloat(7),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorNotEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "bool_not_equal_true",
			left:             ast.NewBool(true),
			right:            ast.NewBool(false),
			operator:         object.OperatorNotEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		{
			name:             "bool_not_equal_false",
			left:             ast.NewBool(true),
			right:            ast.NewBool(true),
			operator:         object.OperatorNotEqual,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		// grater than
		{
			name:             "int_grater_than_false",
			left:             ast.NewNumInt(3),
			right:            ast.NewNumInt(7),
			operator:         object.OperatorGraterThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "int_grater_than_false_equal",
			left:             ast.NewNumInt(7),
			right:            ast.NewNumInt(7),
			operator:         object.OperatorGraterThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "int_grater_than_true",
			left:             ast.NewNumInt(7),
			right:            ast.NewNumInt(3),
			operator:         object.OperatorGraterThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		{
			name:             "float_grater_than_false",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorGraterThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "float_grater_than_false_equal",
			left:             ast.NewNumFloat(7),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorGraterThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjFalse) },
		},
		{
			name:             "float_grater_than_true",
			left:             ast.NewNumFloat(7),
			right:            ast.NewNumFloat(3),
			operator:         object.OperatorGraterThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		// grater or equal than
		{
			name:             "float_grater_or_equal_true",
			left:             ast.NewNumFloat(7),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorGraterOrEqualThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		// less or equal than
		{
			name:             "float_less_or_equal_true",
			left:             ast.NewNumFloat(7),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorLessOrEqualThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
		// less than
		{
			name:             "float_less_than_true",
			left:             ast.NewNumFloat(3),
			right:            ast.NewNumFloat(7),
			operator:         object.OperatorLessThan,
			expectedResultFn: func(r *object.Result) { r.Add(object.ReservedObjTrue) },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comparisonOperation := ast.NewComparisonOperation(tt.left, tt.right, tt.operator)
			env := environment.NewEnvironment()

			expectedRes := object.NewResult()
			tt.expectedResultFn(expectedRes)

			actualRes := object.NewResult()
			execQueue := executor.NewExecFnList()
			packagist := executor.NewPackagist(nil)
			ex := executor.NewExecutor(packagist, execQueue)
			err := comparisonOperation.Exec(env, actualRes, ex)
			require.NoError(t, err, "check error from exec")

			testNextAll(t, execQueue)
			require.NoError(t, err, "check error from fn exec")
			require.Equal(t, expectedRes, actualRes, "check result is equal")
		})
	}
}
