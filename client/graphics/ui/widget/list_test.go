package widget

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestList_SelectedEntry_Initial(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	list := newList(t,
		ListOpts.Entries(entries),

		ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(string)
		}),

		ListOpts.EntrySelectedHandler(func(args *ListEntrySelectedEventArgs) {
			t.Fail() // event fired without previous action
		}))

	assert.Nil(t, list.SelectedEntry())
}

func TestList_SetSelectedEntry(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	var eventArgs *ListEntrySelectedEventArgs
	numEvents := 0

	list := newList(t,
		ListOpts.Entries(entries),

		ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(string)
		}),

		ListOpts.EntrySelectedHandler(func(args *ListEntrySelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	list.SetSelectedEntry(entries[1])
	event.ExecuteDeferred()

	assert.Equal(t, entries[1], eventArgs.Entry)
	assert.Equal(t, entries[1], list.SelectedEntry())

	list.SetSelectedEntry(entries[1])
	event.ExecuteDeferred()

	assert.Equal(t, 1, numEvents)
}

func TestList_EntrySelectedEvent_User(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	var eventArgs *ListEntrySelectedEventArgs
	numEvents := 0

	list := newList(t,
		ListOpts.Entries(entries),

		ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(string)
		}),

		ListOpts.EntrySelectedHandler(func(args *ListEntrySelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	leftMouseButtonClick(list.buttons[1], t)

	assert.Equal(t, entries[1], eventArgs.Entry)
	assert.Equal(t, entries[1], list.SelectedEntry())

	leftMouseButtonClick(list.buttons[1], t)

	assert.Equal(t, 1, numEvents)
}

func TestList_EntrySelectedEvent_User_AllowReselect(t *testing.T) {
	entries := []interface{}{"first", "second", "third"}

	var eventArgs *ListEntrySelectedEventArgs
	numEvents := 0

	list := newList(t,
		ListOpts.Entries(entries),

		ListOpts.EntryLabelFunc(func(e interface{}) string {
			return e.(string)
		}),

		ListOpts.AllowReselect(),

		ListOpts.EntrySelectedHandler(func(args *ListEntrySelectedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	leftMouseButtonClick(list.buttons[1], t)

	assert.Equal(t, entries[1], eventArgs.Entry)
	assert.Equal(t, entries[1], list.SelectedEntry())

	leftMouseButtonClick(list.buttons[1], t)

	assert.Equal(t, entries[1], eventArgs.Entry)
	assert.Equal(t, entries[1], eventArgs.PreviousEntry)
	assert.Equal(t, entries[1], list.SelectedEntry())

	assert.Equal(t, 2, numEvents)
}

func newList(t *testing.T, opts ...ListOpt) *List {
	t.Helper()

	l := NewList(append(opts, []ListOpt{
		ListOpts.ScrollContainerOpts(ScrollContainerOpts.Image(&ScrollContainerImage{
			Idle:     newNineSliceEmpty(t),
			Disabled: newNineSliceEmpty(t),
			Mask:     newNineSliceEmpty(t),
		})),

		ListOpts.SliderOpts(SliderOpts.Images(&SliderTrackImage{}, &ButtonImage{
			Idle: newNineSliceEmpty(t),
		})),

		ListOpts.EntryFontFace(loadFont(t)),

		ListOpts.EntryColor(&ListEntryColor{
			Unselected:                 color.Transparent,
			Selected:                   color.Transparent,
			DisabledUnselected:         color.Transparent,
			DisabledSelected:           color.Transparent,
			SelectedBackground:         color.Transparent,
			DisabledSelectedBackground: color.Transparent,
		}),
	}...)...)

	event.ExecuteDeferred()
	render(l, t)
	return l
}

func listEntryButtons(l *List) []*Button {
	return l.buttons
}
