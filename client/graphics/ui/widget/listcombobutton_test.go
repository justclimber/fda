package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestListComboButton_SelectedEntry_Initial(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	l := newListComboButton(t,
		ListComboButtonOpts.ListOpts(ListOpts.Entries(entries)),

		ListComboButtonOpts.EntrySelectedHandler(func(args *ListComboButtonEntrySelectedEventArgs) {
			t.Fail() // event fired without previous action
		}),

		ListComboButtonOpts.EntryLabelFunc(
			func(e interface{}) string {
				return "label " + e.(string)
			}, func(e interface{}) string {
				return e.(string)
			}))

	assert.Equal(t, entries[0], l.SelectedEntry())
	assert.Equal(t, "label first", l.Label())
}

func TestListComboButton_SetSelectedEntry(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	var eventArgs *ListComboButtonEntrySelectedEventArgs
	numEvents := 0

	l := newListComboButton(t,
		ListComboButtonOpts.ListOpts(ListOpts.Entries(entries)),

		ListComboButtonOpts.EntrySelectedHandler(func(args *ListComboButtonEntrySelectedEventArgs) {
			eventArgs = args
			numEvents++
		}),

		ListComboButtonOpts.EntryLabelFunc(
			func(e interface{}) string {
				return "label " + e.(string)
			}, func(e interface{}) string {
				return e.(string)
			}))

	l.SetSelectedEntry(entries[1])
	event.ExecuteDeferred()

	assert.Equal(t, entries[1], l.SelectedEntry())
	assert.Equal(t, entries[1], eventArgs.Entry)
	assert.Equal(t, entries[0], eventArgs.PreviousEntry)
	assert.Equal(t, "label second", l.Label())

	l.SetSelectedEntry(entries[1])
	event.ExecuteDeferred()
	assert.Equal(t, 1, numEvents)
}

func TestListComboButton_EntrySelectedEvent_User(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	var eventArgs *ListComboButtonEntrySelectedEventArgs
	numEvents := 0

	l := newListComboButton(t,
		ListComboButtonOpts.ListOpts(ListOpts.Entries(entries)),

		ListComboButtonOpts.EntrySelectedHandler(func(args *ListComboButtonEntrySelectedEventArgs) {
			eventArgs = args
			numEvents++
		}),

		ListComboButtonOpts.EntryLabelFunc(
			func(e interface{}) string {
				return "label " + e.(string)
			}, func(e interface{}) string {
				return e.(string)
			}))

	l.SetContentVisible(true)
	render(l, t)

	leftMouseButtonClick(listEntryButtons(listComboButtonContentList(l))[1], t)

	assert.Equal(t, entries[1], l.SelectedEntry())
	assert.Equal(t, entries[1], eventArgs.Entry)
	assert.Equal(t, entries[0], eventArgs.PreviousEntry)
	assert.Equal(t, "label second", l.Label())

	l.SetContentVisible(true)
	render(l, t)

	leftMouseButtonClick(listEntryButtons(listComboButtonContentList(l))[1], t)

	assert.Equal(t, 1, numEvents)
}

func TestListComboButton_ContentVisible_Click(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	l := newListComboButton(t,
		ListComboButtonOpts.ListOpts(ListOpts.Entries(entries)),

		ListComboButtonOpts.EntryLabelFunc(
			func(e interface{}) string {
				return e.(string)
			}, func(e interface{}) string {
				return e.(string)
			}))

	leftMouseButtonClick(l, t)
	assert.True(t, l.ContentVisible())

	leftMouseButtonClick(l, t)
	assert.False(t, l.ContentVisible())
}

func TestListComboButton_ContentVisible_Programmatic(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	l := newListComboButton(t,
		ListComboButtonOpts.ListOpts(ListOpts.Entries(entries)),

		ListComboButtonOpts.EntryLabelFunc(
			func(e interface{}) string {
				return e.(string)
			}, func(e interface{}) string {
				return e.(string)
			}))

	l.SetContentVisible(true)
	assert.True(t, l.ContentVisible())

	l.SetContentVisible(false)
	assert.False(t, l.ContentVisible())
}

func newListComboButton(t *testing.T, opts ...ListComboButtonOpt) *ListComboButton {
	t.Helper()

	l := NewListComboButton(append(opts, []ListComboButtonOpt{
		ListComboButtonOpts.SelectComboButtonOpts(SelectComboButtonOpts.ComboButtonOpts(ComboButtonOpts.ButtonOpts(ButtonOpts.Image(&ButtonImage{
			Idle: newNineSliceEmpty(t),
		})))),
		ListComboButtonOpts.ListOpts(
			ListOpts.ScrollContainerOpts(ScrollContainerOpts.Image(&ScrollContainerImage{
				Idle:     newNineSliceEmpty(t),
				Disabled: newNineSliceEmpty(t),
				Mask:     newNineSliceEmpty(t),
			})),
			ListOpts.SliderOpts(SliderOpts.Images(&SliderTrackImage{}, &ButtonImage{
				Idle: newNineSliceEmpty(t),
			})),
			ListOpts.EntryColor(&ListEntryColor{
				Unselected:                 color.Transparent,
				Selected:                   color.Transparent,
				DisabledUnselected:         color.Transparent,
				DisabledSelected:           color.Transparent,
				SelectedBackground:         color.Transparent,
				DisabledSelectedBackground: color.Transparent,
			}),
			ListOpts.EntryFontFace(loadFont(t)),
		),
		ListComboButtonOpts.Text(loadFont(t), &ButtonImageImage{
			Idle:     newImageEmpty(t),
			Disabled: newImageEmpty(t),
		}, &ButtonTextColor{
			Idle:     color.Transparent,
			Disabled: color.Transparent,
		}),
	}...)...)

	event.ExecuteDeferred()
	render(l, t)
	return l
}

func listComboButtonContentList(l *ListComboButton) *List {
	return l.button.button.content.(*List)
}
