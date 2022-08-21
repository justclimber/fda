package widget

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/justclimber/fda/client/graphics/ui/event"
)

func TestStateButton_SetState_Image(t *testing.T) {
	st := map[interface{}]*ButtonImage{
		1: {Idle: newNineSliceEmpty(t)},
		2: {Idle: newNineSliceEmpty(t)},
		3: {Idle: newNineSliceEmpty(t)},
	}

	s := newStateButton(t, StateButtonOpts.StateImages(st))

	s.State = 2
	render(s, t)

	assert.Equal(t, st[2], stateButtonButton(s).Image)
}

func newStateButton(t *testing.T, opts ...StateButtonOpt) *StateButton {
	s := NewStateButton(opts...)
	event.ExecuteDeferred()
	render(s, t)
	return s
}

func stateButtonButton(s *StateButton) *Button {
	return s.button
}
