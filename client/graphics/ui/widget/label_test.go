package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestLabel_SetLabel(t *testing.T) {
	l := newLabel(t)

	expectedText := "foo"
	l.Label = expectedText
	render(l, t)

	assert.Equal(t, expectedText, labelText(l).Label)
}

func TestLabel_SetDisabled_Color(t *testing.T) {
	l := newLabel(t)

	l.GetWidget().Disabled = true
	render(l, t)

	assert.Equal(t, color.Black, labelText(l).Color)
}

func newLabel(t *testing.T, opts ...LabelOpt) *Label {
	t.Helper()

	l := NewLabel(append(opts, LabelOpts.Text("", loadFont(t), &LabelColor{
		Idle:     color.White,
		Disabled: color.Black,
	}))...)
	event.ExecuteDeferred()
	render(l, t)
	return l
}

func labelText(l *Label) *Text {
	return l.text
}
