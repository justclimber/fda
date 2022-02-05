package integration

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/justclimber/fda/common/configloader"
	"github.com/justclimber/fda/common/debugger"
	"github.com/justclimber/fda/common/debugger/templates"
)

func TestDebugger(t *testing.T) {
	appCfg, err := configloader.NewConfigLoader().Load()
	require.NoError(t, err, "loading config")
	hr := debugger.NewHtmlReport(appCfg.DebuggerHtmlReportFullPath(), templates.EmbeddedFS, time.Second*2)
	d, finish := debugger.NewDebuggerWithReportFinish(true, hr)
	defer finish()
	a := d.CreateNested("A")
	b := d.CreateNested("B")
	c := b.CreateNestedConcurrent("C")
	f := c.CreateNested("F")

	a.LogF("foo", "log 1")
	b.LogF("bar", "log 2")
	c.LogF("foobar", "log 3")
	f.LogF("asd", "log 4")

	assert.False(t, false)
}
