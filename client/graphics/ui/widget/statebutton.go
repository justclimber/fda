package widget

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/justclimber/fda/client/graphics/ui/input"
)

type StateButton struct {
	// TODO: changing this should fire an event
	State interface{}

	buttonOpts []ButtonOpt
	images     map[interface{}]*ButtonImage

	init   *MultiOnce
	button *Button
}

type StateButtonOpt func(s *StateButton)

type StateButtonOptions struct {
}

var StateButtonOpts StateButtonOptions

func NewStateButton(opts ...StateButtonOpt) *StateButton {
	s := &StateButton{
		images: map[interface{}]*ButtonImage{},

		init: &MultiOnce{},
	}

	s.init.Append(s.createWidget)

	for _, o := range opts {
		o(s)
	}

	return s
}

func (o StateButtonOptions) ButtonOpts(opts ...ButtonOpt) StateButtonOpt {
	return func(s *StateButton) {
		s.buttonOpts = append(s.buttonOpts, opts...)
	}
}

func (o StateButtonOptions) StateImages(states map[interface{}]*ButtonImage) StateButtonOpt {
	return func(s *StateButton) {
		initial := true
		for st, i := range states {
			s.images[st] = i

			if initial {
				s.State = st
				initial = false
			}
		}
	}
}

func (s *StateButton) GetWidget() *Widget {
	s.init.Do()
	return s.button.GetWidget()
}

func (s *StateButton) PreferredSize() (int, int) {
	s.init.Do()
	return s.button.PreferredSize()
}

func (s *StateButton) SetLocation(rect image.Rectangle) {
	s.init.Do()
	s.button.SetLocation(rect)
}

func (s *StateButton) RequestRelayout() {
	s.init.Do()
	s.button.RequestRelayout()
}

func (s *StateButton) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	s.init.Do()
	s.button.SetupInputLayer(def)
}

func (s *StateButton) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	s.init.Do()

	s.button.Image = s.images[s.State]

	s.button.Render(screen, def, debugMode)
}

func (s *StateButton) createWidget() {
	s.button = NewButton(append(s.buttonOpts, ButtonOpts.Image(s.images[s.State]))...)
	s.buttonOpts = nil
}
