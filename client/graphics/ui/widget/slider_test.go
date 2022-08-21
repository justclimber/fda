package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestSlider_Current_Initial(t *testing.T) {
	var eventArgs *SliderChangedEventArgs
	s := newSlider(t,
		SliderOpts.MinMax(10, 20),
		SliderOpts.ChangedHandler(func(args *SliderChangedEventArgs) {
			eventArgs = args
		}))

	assert.Equal(t, 10, s.Current)
	assert.Equal(t, 10, eventArgs.Current)
}

func newSlider(t *testing.T, opts ...SliderOpt) *Slider {
	s := NewSlider(append(opts, SliderOpts.Images(&SliderTrackImage{
		Idle: newNineSliceEmpty(t),
	}, &ButtonImage{
		Idle: newNineSliceEmpty(t),
	}))...)
	event.ExecuteDeferred()
	render(s, t)
	return s
}
