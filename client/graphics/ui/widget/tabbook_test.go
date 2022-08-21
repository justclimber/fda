package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestTabBook_Tab_Initial(t *testing.T) {
	tab1 := NewTabBookTab("Tab 1", newSimpleWidget(50, 50, nil))
	tab2 := NewTabBookTab("Tab 2", newSimpleWidget(50, 50, nil))

	tb := newTabBook(t,
		TabBookOpts.Tabs(tab1, tab2),
		TabBookOpts.TabSelectedHandler(func(args *TabBookTabSelectedEventArgs) {
			t.Fail() // event fired without previous action
		}))

	assert.Equal(t, tab1, tb.Tab())
}

func TestTabBook_SetTab(t *testing.T) {
	var eventArgs *TabBookTabSelectedEventArgs
	numEvents := 0

	tab1 := NewTabBookTab("Tab 1", newSimpleWidget(50, 50, nil))
	tab2 := NewTabBookTab("Tab 2", newSimpleWidget(50, 50, nil))

	tb := newTabBook(t,
		TabBookOpts.Tabs(tab1, tab2),
		TabBookOpts.TabSelectedHandler(func(args *TabBookTabSelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	tb.SetTab(tab2)
	event.ExecuteDeferred()

	assert.Equal(t, tab2, tb.Tab())
	assert.Equal(t, tab2, eventArgs.Tab)
	assert.Equal(t, tab1, eventArgs.PreviousTab)

	tb.SetTab(tab2)
	event.ExecuteDeferred()
	assert.Equal(t, 1, numEvents)
}

func TestTabBook_TabSelectedEvent_User(t *testing.T) {
	var eventArgs *TabBookTabSelectedEventArgs
	numEvents := 0

	tab1 := NewTabBookTab("Tab 1", newSimpleWidget(50, 50, nil))
	tab2 := NewTabBookTab("Tab 2", newSimpleWidget(50, 50, nil))

	tb := newTabBook(t,
		TabBookOpts.Tabs(tab1, tab2),
		TabBookOpts.TabSelectedHandler(func(args *TabBookTabSelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	leftMouseButtonClick(tabBookButtons(tb)[1], t)

	assert.Equal(t, tab2, tb.Tab())
	assert.Equal(t, tab2, eventArgs.Tab)
	assert.Equal(t, tab1, eventArgs.PreviousTab)

	leftMouseButtonClick(tabBookButtons(tb)[1], t)
	assert.Equal(t, 1, numEvents)
}

func newTabBook(t *testing.T, opts ...TabBookOpt) *TabBook {
	t.Helper()

	tb := NewTabBook(append(opts, []TabBookOpt{
		TabBookOpts.TabButtonImage(&ButtonImage{
			Idle: newNineSliceEmpty(t),
		}, &ButtonImage{
			Idle: newNineSliceEmpty(t),
		}),
		TabBookOpts.TabButtonText(loadFont(t), &ButtonTextColor{
			Idle:     color.Transparent,
			Disabled: color.Transparent,
		}),
	}...)...)

	event.ExecuteDeferred()
	render(tb, t)
	return tb
}

func tabBookButtons(t *TabBook) []*StateButton {
	buttons := []*StateButton{}
	for _, tab := range t.tabs {
		buttons = append(buttons, t.tabToButton[tab])
	}
	return buttons
}
