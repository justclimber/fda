package widget

import (
	"image"
	img "image"
	"sync/atomic"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/justclimber/fda/client/graphics/ui/input"
)

type ToolTip struct {
	Sticky bool
	Delay  time.Duration

	container       Locater
	contentsCreater ToolTipContentsCreater
	offset          img.Point

	state toolTipState
}

type ToolTipOpt func(t *ToolTip)

type toolTipState func(*ebiten.Image, DeferredRenderFunc) (toolTipState, bool)

type ToolTipContentsCreater interface {
	Create(HasWidget) ToolTipWidget
}

type ToolTipWidget interface {
	PreferredSizer
	Locateable
	Renderer
}

type ToolTipContentsUpdater interface {
	Update(HasWidget)
}

type ToolTipOptions struct {
}

var ToolTipOpts ToolTipOptions

func NewToolTip(opts ...ToolTipOpt) *ToolTip {
	t := &ToolTip{
		offset: img.Point{0, 20},
	}
	t.state = t.idleState()

	for _, o := range opts {
		o(t)
	}

	return t
}

func (o ToolTipOptions) Container(c Locater) ToolTipOpt {
	return func(t *ToolTip) {
		t.container = c
	}
}

func (o ToolTipOptions) ContentsCreater(c ToolTipContentsCreater) ToolTipOpt {
	return func(t *ToolTip) {
		t.contentsCreater = c
	}
}

func (o ToolTipOptions) Offset(off img.Point) ToolTipOpt {
	return func(t *ToolTip) {
		t.offset = off
	}
}

func (o ToolTipOptions) Sticky() ToolTipOpt {
	return func(t *ToolTip) {
		t.Sticky = true
	}
}

func (o ToolTipOptions) Delay(d time.Duration) ToolTipOpt {
	return func(t *ToolTip) {
		t.Delay = d
	}
}

func (t *ToolTip) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	for {
		newState, rerun := t.state(screen, def)
		if newState != nil {
			t.state = newState
		}
		if !rerun {
			break
		}
	}
}

func (t *ToolTip) idleState() toolTipState {
	return func(screen *ebiten.Image, def DeferredRenderFunc) (toolTipState, bool) {
		if input.MouseButtonPressed(ebiten.MouseButtonLeft) ||
			input.MouseButtonPressed(ebiten.MouseButtonMiddle) ||
			input.MouseButtonPressed(ebiten.MouseButtonRight) {

			return nil, false
		}

		x, y := input.CursorPosition()
		w := t.container.WidgetAt(x, y)
		if w == nil {
			return nil, false
		}

		if t.Delay <= 0 {
			return t.showingState(w, x, y, nil), true
		}

		return t.armedState(w, x, y, nil, nil), true
	}
}

func (t *ToolTip) armedState(srcWidget HasWidget, srcX int, srcY int, timer *time.Timer, expired *atomic.Value) toolTipState {
	return func(screen *ebiten.Image, def DeferredRenderFunc) (toolTipState, bool) {
		x, y := input.CursorPosition()
		w := t.container.WidgetAt(x, y)
		if input.MouseButtonPressed(ebiten.MouseButtonLeft) ||
			input.MouseButtonPressed(ebiten.MouseButtonMiddle) ||
			input.MouseButtonPressed(ebiten.MouseButtonRight) ||
			w != srcWidget {

			return t.idleState(), false
		}

		if timer != nil && expired.Load().(bool) {
			return t.showingState(srcWidget, x, y, nil), true
		}

		if timer == nil {
			expired = &atomic.Value{}
			expired.Store(false)

			timer = time.AfterFunc(t.Delay, func() {
				expired.Store(true)
			})

			return t.armedState(srcWidget, srcX, srcY, timer, expired), false
		}

		return nil, false
	}
}

func (t *ToolTip) showingState(srcWidget HasWidget, srcX int, srcY int, tipWidget ToolTipWidget) toolTipState {
	return func(screen *ebiten.Image, def DeferredRenderFunc) (toolTipState, bool) {
		x, y := input.CursorPosition()
		w := t.container.WidgetAt(x, y)
		if input.MouseButtonPressed(ebiten.MouseButtonLeft) ||
			input.MouseButtonPressed(ebiten.MouseButtonMiddle) ||
			input.MouseButtonPressed(ebiten.MouseButtonRight) ||
			w != srcWidget {

			return t.idleState(), false
		}

		if tipWidget == nil {
			tipWidget = t.contentsCreater.Create(srcWidget)

			if tipWidget == nil {
				return t.idleState(), false
			}
		}

		if u, ok := t.contentsCreater.(ToolTipContentsUpdater); ok {
			u.Update(srcWidget)
		}

		if !t.Sticky {
			srcX, srcY = x, y
		}

		sx, sy := tipWidget.PreferredSize()
		r := image.Rect(0, 0, sx, sy)
		r = r.Add(image.Point{srcX, srcY})
		r = r.Add(t.offset)
		tipWidget.SetLocation(r)
		if rl, ok := tipWidget.(Relayoutable); ok {
			rl.RequestRelayout()
		}
		tipWidget.Render(screen, def, DebugModeNone)

		return t.showingState(srcWidget, srcX, srcY, tipWidget), false
	}
}
