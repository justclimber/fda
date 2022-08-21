package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestCheckbox_State_Initial(t *testing.T) {
	c := newCheckbox(t,
		CheckboxOpts.ChangedHandler(func(args *CheckboxChangedEventArgs) {
			t.Fail() // event fired without previous action
		}))

	assert.Equal(t, CheckboxUnchecked, c.State())
}

func TestCheckbox_ChangedEvent_User(t *testing.T) {
	var eventArgs *CheckboxChangedEventArgs

	c := newCheckbox(t,
		CheckboxOpts.ChangedHandler(func(args *CheckboxChangedEventArgs) {
			eventArgs = args
		}))

	leftMouseButtonClick(c, t)

	assert.Equal(t, CheckboxChecked, eventArgs.State)
	assert.Equal(t, CheckboxChecked, c.State())
}

func TestCheckbox_SetState(t *testing.T) {
	var eventArgs *CheckboxChangedEventArgs
	numEvents := 0

	c := newCheckbox(t,
		CheckboxOpts.ChangedHandler(func(args *CheckboxChangedEventArgs) {
			eventArgs = args
			numEvents++
		}))

	c.SetState(CheckboxChecked)
	event.ExecuteDeferred()

	assert.Equal(t, CheckboxChecked, eventArgs.State)
	assert.Equal(t, CheckboxChecked, c.State())

	c.SetState(CheckboxChecked)
	event.ExecuteDeferred()

	assert.Equal(t, 1, numEvents)
}

func TestCheckbox_State_Cycle(t *testing.T) {
	c := newCheckbox(t)
	leftMouseButtonClick(c, t)
	assert.Equal(t, CheckboxChecked, c.State())
	leftMouseButtonClick(c, t)
	assert.Equal(t, CheckboxUnchecked, c.State())
}

func TestCheckbox_State_Cycle_TriState(t *testing.T) {
	c := newCheckbox(t, CheckboxOpts.TriState())
	leftMouseButtonClick(c, t)
	assert.Equal(t, CheckboxChecked, c.State())
	leftMouseButtonClick(c, t)
	assert.Equal(t, CheckboxGreyed, c.State())
	leftMouseButtonClick(c, t)
	assert.Equal(t, CheckboxUnchecked, c.State())
}

func newCheckbox(t *testing.T, opts ...CheckboxOpt) *Checkbox {
	t.Helper()

	c := NewCheckbox(append(opts, []CheckboxOpt{
		CheckboxOpts.ButtonOpts(ButtonOpts.Image(&ButtonImage{
			Idle: newNineSliceEmpty(t),
		})),

		CheckboxOpts.Image(&CheckboxGraphicImage{
			Unchecked: &ButtonImageImage{
				Idle: newImageEmpty(t),
			},
			Checked: &ButtonImageImage{
				Idle: newImageEmpty(t),
			},
		}),
	}...)...)
	event.ExecuteDeferred()
	render(c, t)
	return c
}
