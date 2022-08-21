package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestRadioGroup_Active_Initial(t *testing.T) {
	var cbs []*Checkbox
	for i := 0; i < 3; i++ {
		c := newCheckbox(t)
		cbs = append(cbs, c)
	}

	var eventArgs *RadioGroupChangedEventArgs

	r := newRadioGroup(t, cbs, RadioGroupOpts.ChangedHandler(func(args *RadioGroupChangedEventArgs) {
		eventArgs = args
	}))

	assert.Equal(t, cbs[0], r.Active())
	assert.Equal(t, cbs[0], eventArgs.Active)
	assert.Equal(t, CheckboxChecked, cbs[0].State())
}

func TestRadioGroup_ChangedEvent_User(t *testing.T) {
	var cbs []*Checkbox
	for i := 0; i < 3; i++ {
		c := newCheckbox(t)
		cbs = append(cbs, c)
	}

	r := newRadioGroup(t, cbs)

	var eventArgs *RadioGroupChangedEventArgs
	r.ChangedEvent.AddHandler(func(args interface{}) {
		eventArgs = args.(*RadioGroupChangedEventArgs)
	})

	leftMouseButtonClick(cbs[1], t)

	assert.Equal(t, cbs[1], r.Active())
	assert.Equal(t, cbs[1], eventArgs.Active)
	assert.Equal(t, CheckboxUnchecked, cbs[0].State())
	assert.Equal(t, CheckboxChecked, cbs[1].State())
}

func TestRadioGroup_SetActive(t *testing.T) {
	var cbs []*Checkbox
	for i := 0; i < 3; i++ {
		c := newCheckbox(t)
		cbs = append(cbs, c)
	}

	r := newRadioGroup(t, cbs)

	var eventArgs *RadioGroupChangedEventArgs
	r.ChangedEvent.AddHandler(func(args interface{}) {
		eventArgs = args.(*RadioGroupChangedEventArgs)
	})

	r.SetActive(cbs[1])
	event.ExecuteDeferred()

	assert.Equal(t, cbs[1], r.Active())
	assert.Equal(t, cbs[1], eventArgs.Active)
	assert.Equal(t, CheckboxUnchecked, cbs[0].State())
	assert.Equal(t, CheckboxChecked, cbs[1].State())
}

func newRadioGroup(t *testing.T, cbs []*Checkbox, opts ...RadioGroupOpt) *RadioGroup {
	t.Helper()

	r := NewRadioGroup(append(opts, RadioGroupOpts.Checkboxes(cbs...))...)
	event.ExecuteDeferred()
	for _, c := range cbs {
		render(c, t)
	}
	return r
}
