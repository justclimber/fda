package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestSelectComboButton_SetSelectedEntry(t *testing.T) {
	var eventArgs *SelectComboButtonEntrySelectedEventArgs
	numEvents := 0

	b := newSelectComboButton(t,
		SelectComboButtonOpts.EntryLabelFunc(func(e interface{}) string {
			return "label " + e.(string)
		}),

		SelectComboButtonOpts.EntrySelectedHandler(func(args *SelectComboButtonEntrySelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	entry := "foo"
	b.SetSelectedEntry(entry)
	event.ExecuteDeferred()

	assert.Equal(t, entry, b.SelectedEntry())
	assert.Equal(t, entry, eventArgs.Entry)
	assert.Equal(t, "label foo", b.Label())

	b.SetSelectedEntry(entry)
	event.ExecuteDeferred()

	assert.Equal(t, 1, numEvents)

	entry2 := "bar"
	b.SetSelectedEntry(entry2)
	event.ExecuteDeferred()

	assert.Equal(t, entry, eventArgs.PreviousEntry)
}

func TestSelectComboButton_ContentVisible_Click(t *testing.T) {
	b := newSelectComboButton(t)

	leftMouseButtonClick(b, t)
	assert.True(t, b.ContentVisible())

	leftMouseButtonClick(b, t)
	assert.False(t, b.ContentVisible())
}

func TestSelectComboButton_ContentVisible_Programmatic(t *testing.T) {
	b := newSelectComboButton(t)

	b.SetContentVisible(true)
	event.ExecuteDeferred()

	assert.True(t, b.ContentVisible())

	b.SetContentVisible(false)
	event.ExecuteDeferred()

	assert.False(t, b.ContentVisible())
}

func newSelectComboButton(t *testing.T, opts ...SelectComboButtonOpt) *SelectComboButton {
	t.Helper()

	b := NewSelectComboButton(append(opts,
		SelectComboButtonOpts.ComboButtonOpts(
			ComboButtonOpts.ButtonOpts(
				ButtonOpts.Image(&ButtonImage{
					Idle: newNineSliceEmpty(t),
				}),
				ButtonOpts.TextAndImage("", loadFont(t), &ButtonImageImage{
					Idle:     newImageEmpty(t),
					Disabled: newImageEmpty(t),
				}, &ButtonTextColor{
					Idle:     color.Transparent,
					Disabled: color.Transparent,
				}),
			),
			ComboButtonOpts.Content(newButton(t))),
	)...)

	event.ExecuteDeferred()
	render(b, t)
	return b
}
