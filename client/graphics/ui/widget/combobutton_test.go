package widget

import (
	"testing"

	"github.com/matryer/is"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestComboButton_ContentVisible_Click(t *testing.T) {
	is := is.New(t)

	b := newComboButton(t)

	leftMouseButtonClick(b, t)
	is.True(b.ContentVisible)

	leftMouseButtonClick(b, t)
	is.True(!b.ContentVisible)
}

func newComboButton(t *testing.T, opts ...ComboButtonOpt) *ComboButton {
	t.Helper()

	b := NewComboButton(append(opts, []ComboButtonOpt{
		ComboButtonOpts.ButtonOpts(ButtonOpts.Image(&ButtonImage{
			Idle: newNineSliceEmpty(t),
		})),
		ComboButtonOpts.Content(newButton(t)),
	}...)...)
	event.ExecuteDeferred()
	render(b, t)
	return b
}
