package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestComboButton_ContentVisible_Click(t *testing.T) {

	b := newComboButton(t)

	leftMouseButtonClick(b, t)
	assert.True(t, b.ContentVisible)

	leftMouseButtonClick(b, t)
	assert.False(t, b.ContentVisible)
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
