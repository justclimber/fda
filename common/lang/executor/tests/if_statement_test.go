package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
)

func TestIfStatement_Exec_WithoutFalseBranch(t *testing.T) {
	tests := []struct {
		name          string
		conditionExpr ast.Expr
		checkVar      bool
	}{
		{
			name:          "check_true_branch",
			conditionExpr: ast.NewBool(true),
			checkVar:      true,
		},
		{
			name:          "false_branch",
			conditionExpr: ast.NewBool(false),
			checkVar:      false,
		},
	}
	expectedInt := int64(44)
	varName := "a"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifStmtAst := ast.NewIfStatement(
				tt.conditionExpr,
				ast.NewStatementsBlock(0, []ast.Stmt{
					ast.NewVoidedExpression(0, ast.NewAssignment(
						0,
						[]*ast.Identifier{ast.NewIdentifier(0, varName)},
						ast.NewNumInt(0, expectedInt),
					)),
				}),
				nil,
			)
			env := environment.NewEnvironment()
			packagist := executor.NewPackagist(nil)
			execQueue := executor.NewExecFnList()
			ex := executor.NewExecutor(packagist, execQueue)
			err := ifStmtAst.Exec(env, ex)
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
