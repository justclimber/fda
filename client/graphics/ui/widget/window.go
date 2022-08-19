package widget

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/justclimber/fda/client/graphics/ui/input"
)

type Window struct {
	ID    int
	Modal bool
	state windowState

	init *MultiOnce

	movable   *Container
	contents  *Container
	container *Container
}

type windowState func(*ebiten.Image, DeferredRenderFunc) (nextState windowState, rerun bool)

type WindowOpt func(w *Window)

type WindowOptions struct {
}

var WindowOpts WindowOptions

func NewWindow(opts ...WindowOpt) *Window {
	w := &Window{
		init: &MultiOnce{},
	}

	for _, o := range opts {
		o(w)
	}
	w.init.Append(w.createWidget)

	return w
}

func (o WindowOptions) Contents(c *Container) WindowOpt {
	return func(w *Window) {
		w.contents = c
	}
}

func (o WindowOptions) Movable(c *Container) WindowOpt {
	return func(w *Window) {
		w.movable = c
	}
}

func (o WindowOptions) Modal() WindowOpt {
	return func(w *Window) {
		w.Modal = true
	}
}

func (w *Window) createWidget() {
	if w.movable != nil {
		w.container = NewContainer(
			"window container",
			ContainerOpts.Layout(NewGridLayout(
				GridLayoutOpts.Stretch([]bool{true}, []bool{false, true}),
				GridLayoutOpts.Columns(1),
			)),
		)
		w.container.AddChild(w.movable)
		w.state = w.idleState()
		w.container.AddChild(w.contents)
	} else {
		w.container = w.contents
	}
}

func (w *Window) Container() *Container {
	w.init.Do()
	return w.container
}

func (w *Window) SetLocation(rect image.Rectangle) {
	w.init.Do()
	w.container.SetLocation(rect)
}

func (w *Window) RequestRelayout() {
	w.init.Do()
	w.container.RequestRelayout()
}

func (w *Window) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	w.init.Do()
	var l *input.Layer
	if w.Modal {
		l = &input.Layer{
			DebugLabel: "modal window",
			EventTypes: input.LayerEventTypeAll,
			BlockLower: true,
			FullScreen: true,
		}
	} else {
		l = &input.Layer{
			DebugLabel: "window",
			EventTypes: input.LayerEventTypeAll,
			BlockLower: true,
			RectFunc: func() image.Rectangle {
				return w.container.GetWidget().Rect
			},
		}
	}
	w.container.GetWidget().ElevateToNewInputLayer(l)

	if w.movable != nil {
		w.movable.GetWidget().ElevateToNewInputLayer(&input.Layer{
			DebugLabel: "window movable",
			EventTypes: input.LayerEventTypeMouseButton,
			BlockLower: true,
			RectFunc: func() image.Rectangle {
				return w.movable.GetWidget().Rect
			},
		})
	}
	w.container.SetupInputLayer(def)
}

func (w *Window) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	w.init.Do()
	w.runState(screen, def)
	w.container.Render(screen, def, debugMode)
}

func (w *Window) runState(screen *ebiten.Image, def DeferredRenderFunc) {
	if w.state != nil {
		for {
			newState, rerun := w.state(screen, def)
			if newState != nil {
				w.state = newState
			}
			if !rerun {
				break
			}
		}
	}
}

func (w *Window) idleState() windowState {
	return func(screen *ebiten.Image, def DeferredRenderFunc) (windowState, bool) {
		if !input.MouseButtonJustPressedLayer(ebiten.MouseButtonLeft, w.movable.widget.EffectiveInputLayer()) {
			return nil, false
		}

		x, y := input.CursorPosition()
		return w.dragState(x, y), false
	}
}

func (w *Window) dragState(oldX int, oldY int) windowState {
	return func(screen *ebiten.Image, def DeferredRenderFunc) (windowState, bool) {
		if !input.MouseButtonPressed(ebiten.MouseButtonLeft) {
			return w.idleState(), false
		}
		x, y := input.CursorPosition()
		dx := x - oldX
		dy := y - oldY
		if dx != 0 || dy != 0 {
			rect := w.container.widget.Rect
			rect = rect.Add(image.Point{X: dx, Y: dy})
			w.SetLocation(rect)
			w.RequestRelayout()
		}

		return w.dragState(x, y), false
	}
}
