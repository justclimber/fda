package executor

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/fdalang"
)

func TestIfStatement_Exec_WithoutFalseBranch(t *testing.T) {
	tests := []struct {
		name          string
		conditionExpr Expr
		checkVar      bool
	}{
		{
			name:          "check_true_branch",
			conditionExpr: NewBool(true),
			checkVar:      true,
		},
		{
			name:          "false_branch",
			conditionExpr: NewBool(false),
			checkVar:      false,
		},
	}
	expectedInt := int64(44)
	varName := "a"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ast := NewIfStatement(
				tt.conditionExpr,
				NewStatementsBlock([]Stmt{
					NewVoidedExpression(
						NewAssignment(
							NewIdentifierList([]string{varName}),
							NewNumInt(expectedInt),
						),
					),
				}),
				nil,
			)
			env := fdalang.NewEnvironment()
			execQueue := NewExecFnList()
			err := ast.Exec(env, execQueue)
			require.NoError(t, err, "check error from exec")

			testNextAll(t, execQueue)
			obj, ok := env.Get(varName)
			require.Equal(t, tt.checkVar, ok, "check existence var in env")
			if tt.checkVar {
				testObjectAsNumInt(t, obj, expectedInt)
			}
		})
	}
}
