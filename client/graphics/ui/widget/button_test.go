package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestButton_PressedEvent_User(t *testing.T) {
	var eventArgs *ButtonPressedEventArgs

	b := newButton(t,
		ButtonOpts.PressedHandler(func(args *ButtonPressedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonPress(b, t)
	assert.NotNil(t, eventArgs)
}

func TestButton_ReleasedEvent_User(t *testing.T) {
	var eventArgs *ButtonReleasedEventArgs

	b := newButton(t,
		ButtonOpts.ReleasedHandler(func(args *ButtonReleasedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonRelease(b, t)
	assert.NotNil(t, eventArgs)
}

func TestButton_ClickedEvent_User(t *testing.T) {
	var eventArgs *ButtonClickedEventArgs

	b := newButton(t,
		ButtonOpts.ClickedHandler(func(args *ButtonClickedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonClick(b, t)
	assert.NotNil(t, eventArgs)
}

func newButton(t *testing.T, opts ...ButtonOpt) *Button {
	t.Helper()

	b := NewButton(append(opts, ButtonOpts.Image(&ButtonImage{
		Idle: newNineSliceEmpty(t),
	}))...)
	event.ExecuteDeferred()
	render(b, t)
	return b
}
