package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
	"github.com/justclimber/fda/common/lang/executor/object"
)

func testNextAll(t *testing.T, execQueue *executor.ExecFnList) {
	t.Helper()
	var err error
	hasNext := true
	for hasNext {
		hasNext, err = execQueue.ExecNext()
		require.NoError(t, err, "check error from fn exec")
	}
}

func testObjectAsNumInt(t *testing.T, obj object.Object, expectedInt int64) {
	t.Helper()
	objInt, ok := obj.(*object.ObjInteger)
	require.True(t, ok, "check obj type")

	assert.Equal(t, expectedInt, objInt.Value)
}
