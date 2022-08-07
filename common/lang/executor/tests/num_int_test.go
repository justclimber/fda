package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/ast"
	"github.com/justclimber/fda/common/lang/executor/environment"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func TestNumInt_Exec(t *testing.T) {
	expectedInt := int64(4)
	numInt := ast.NewNumInt(0, expectedInt)

	env := environment.NewEnvironment()
	res := object.NewResult()
	packagist := executor.NewPackagist(nil)
	execQueue := executor.NewExecFnList()
	ex := executor.NewExecutor(packagist, execQueue)
	err := numInt.Exec(env, res, ex)
	require.NoError(t, err, "check error from exec")

	testNextAll(t, execQueue)
	require.NoError(t, err, "check error from fn exec")
	testResultAsNumInt(t, res, expectedInt, 0)
}
