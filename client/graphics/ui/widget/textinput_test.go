package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestTextInput_ChangedEvent(t *testing.T) {
	var eventArgs *TextInputChangedEventArgs
	ti := newTextInput(t, TextInputOpts.ChangedHandler(func(args *TextInputChangedEventArgs) {
		eventArgs = args
	}))

	ti.InputText = "foo"
	render(ti, t)

	assert.Equal(t, "foo", eventArgs.InputText)
}

func TestTextInput_ChangedEvent_OnlyOnce(t *testing.T) {
	numEvents := 0
	ti := newTextInput(t, TextInputOpts.ChangedHandler(func(args *TextInputChangedEventArgs) {
		numEvents++
	}))

	ti.InputText = "foo"
	render(ti, t)
	render(ti, t)

	assert.Equal(t, 1, numEvents)
}

func TestTextInput_DoBackspace(t *testing.T) {
	ti := newTextInput(t)
	ti.InputText = "foo"
	ti.cursorPosition = 1
	render(ti, t)

	ti.ChangedEvent.AddHandler(func(args interface{}) {
		assert.Equal(t, "oo", args.(*TextInputChangedEventArgs).InputText)
	})

	ti.doBackspace()
	render(ti, t)
}

func TestTextInput_DoBackspace_Disabled(t *testing.T) {
	ti := newTextInput(t)
	ti.GetWidget().Disabled = true
	ti.InputText = "foo"
	ti.cursorPosition = 1
	render(ti, t)

	ti.ChangedEvent.AddHandler(func(args interface{}) {
		t.Fail() // received event even though widget is disabled
	})

	ti.doBackspace()
	render(ti, t)
}

func TestTextInput_DoDelete(t *testing.T) {
	ti := newTextInput(t)
	ti.InputText = "foo"
	render(ti, t)

	ti.ChangedEvent.AddHandler(func(args interface{}) {
		assert.Equal(t, "oo", args.(*TextInputChangedEventArgs).InputText)
	})

	ti.doDelete()
	render(ti, t)
}

func TestTextInput_DoDelete_Disabled(t *testing.T) {
	ti := newTextInput(t)
	ti.GetWidget().Disabled = true
	ti.InputText = "foo"
	render(ti, t)

	ti.ChangedEvent.AddHandler(func(args interface{}) {
		t.Fail() // received event even though widget is disabled
	})

	ti.doDelete()
	render(ti, t)
}

func TestTextInput_DoInsert(t *testing.T) {
	ti := newTextInput(t)
	ti.InputText = "foo"
	ti.cursorPosition = 1
	render(ti, t)

	ti.doInsert([]rune("ab€c"))

	assert.Equal(t, "fab€coo", ti.InputText)
	assert.Equal(t, 5, ti.cursorPosition)
}

func newTextInput(t *testing.T, opts ...TextInputOpt) *TextInput {
	ti := NewTextInput(append(opts, []TextInputOpt{
		TextInputOpts.Face(loadFont(t)),
		TextInputOpts.Color(&TextInputColor{
			Idle:     color.White,
			Disabled: color.White,
			Caret:    color.White,
		}),
		TextInputOpts.CaretOpts(
			CaretOpts.Size(loadFont(t), 1)),
	}...)...)
	event.ExecuteDeferred()
	render(ti, t)
	return ti
}
