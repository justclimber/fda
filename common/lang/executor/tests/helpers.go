package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/lang/executor"
)

func testNext(t *testing.T, execQueue *executor.ExecFnList, times int) {
	t.Helper()
	for i := 0; i < times; i++ {
		_, err := execQueue.ExecNext()
		require.NoError(t, err, "check error from fn exec")
	}
}

func testNextAll(t *testing.T, execQueue *executor.ExecFnList) {
	t.Helper()
	var err error
	hasNext := true
	for hasNext {
		hasNext, err = execQueue.ExecNext()
		require.NoError(t, err, "check error from fn exec")
	}
}
