package widget

import (
	img "image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/justclimber/fda/client/graphics/ui/event"
	"github.com/justclimber/fda/client/graphics/ui/image"
	"github.com/justclimber/fda/client/graphics/ui/input"
)

type Slider struct {
	Min               int
	Max               int
	Current           int
	DrawTrackDisabled bool

	ChangedEvent *event.Event

	widgetOpts   []WidgetOpt
	handleOpts   []ButtonOpt
	direction    Direction
	trackImage   *SliderTrackImage
	trackPadding Insets
	handleSize   int
	pageSizeFunc SliderPageSizeFunc

	init                         *MultiOnce
	widget                       *Widget
	handle                       *Button
	lastCurrent                  int
	hovering                     bool
	dragging                     bool
	handlePressedCursorX         int
	handlePressedCursorY         int
	handlePressedOffsetX         int
	handlePressedOffsetY         int
	handlePressedInternalCurrent float64
}

type SliderTrackImage struct {
	Idle     *image.NineSlice
	Hover    *image.NineSlice
	Disabled *image.NineSlice
}

type SliderOpt func(s *Slider)

type SliderPageSizeFunc func() int

type SliderChangedEventArgs struct {
	Slider   *Slider
	Current  int
	Dragging bool
}

type SliderChangedHandlerFunc func(args *SliderChangedEventArgs)

type SliderOptions struct {
}

var SliderOpts SliderOptions

func NewSlider(opts ...SliderOpt) *Slider {
	s := &Slider{
		Min:     1,
		Max:     100,
		Current: 1,

		ChangedEvent: &event.Event{},

		trackImage: &SliderTrackImage{},
		handleSize: 16,
		pageSizeFunc: func() int {
			return 10
		},

		lastCurrent: 1,

		init: &MultiOnce{},
	}

	s.init.Append(s.createWidget)

	for _, o := range opts {
		o(s)
	}

	return s
}

func (o SliderOptions) WidgetOpts(opts ...WidgetOpt) SliderOpt {
	return func(s *Slider) {
		s.widgetOpts = append(s.widgetOpts, opts...)
	}
}

func (o SliderOptions) Direction(d Direction) SliderOpt {
	return func(s *Slider) {
		s.direction = d
	}
}

func (o SliderOptions) Images(track *SliderTrackImage, handle *ButtonImage) SliderOpt {
	return func(s *Slider) {
		s.trackImage = track
		s.handleOpts = append(s.handleOpts, ButtonOpts.Image(handle))
	}
}

func (o SliderOptions) TrackPadding(i Insets) SliderOpt {
	return func(s *Slider) {
		s.trackPadding = i
	}
}

func (o SliderOptions) HandleSize(s int) SliderOpt {
	return func(sl *Slider) {
		sl.handleSize = s
	}
}

func (o SliderOptions) MinMax(min int, max int) SliderOpt {
	return func(s *Slider) {
		s.Min = min
		s.Max = max
	}
}

func (o SliderOptions) PageSizeFunc(f SliderPageSizeFunc) SliderOpt {
	return func(s *Slider) {
		s.pageSizeFunc = f
	}
}

func (o SliderOptions) ChangedHandler(f SliderChangedHandlerFunc) SliderOpt {
	return func(s *Slider) {
		s.ChangedEvent.AddHandler(func(args interface{}) {
			f(args.(*SliderChangedEventArgs))
		})
	}
}

func (s *Slider) GetWidget() *Widget {
	s.init.Do()
	return s.widget
}

func (s *Slider) PreferredSize() (int, int) {
	if s.direction == DirectionHorizontal {
		return 200, s.handleSize + s.trackPadding.Top + s.trackPadding.Bottom
	}

	return s.handleSize + s.trackPadding.Left + s.trackPadding.Right, 200
}

func (s *Slider) SetLocation(rect img.Rectangle) {
	s.init.Do()
	s.widget.Rect = rect
}

func (s *Slider) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	s.handle.GetWidget().ElevateToNewInputLayer(&input.Layer{
		DebugLabel: "slider handle",
		EventTypes: input.LayerEventTypeAll,
		BlockLower: true,
		FullScreen: false,
		RectFunc: func() img.Rectangle {
			return s.handle.GetWidget().Rect
		},
	})

	s.handle.SetupInputLayer(def)
}

func (s *Slider) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	s.init.Do()

	s.clampCurrentMinMax()
	s.handle.GetWidget().Disabled = s.widget.Disabled

	s.widget.Render(screen, def, debugMode)

	s.draw(screen)

	hl, tl := s.handleLengthAndTrackLength()
	s.updateHandleLocation(hl, tl)
	s.updateHandleSize(hl)

	s.handle.Render(screen, def, debugMode)

	s.fireEvents()

	s.lastCurrent = s.Current
}

func (s *Slider) draw(screen *ebiten.Image) {
	i := s.trackImage.Idle
	if s.widget.Disabled || s.DrawTrackDisabled {
		if s.trackImage.Disabled != nil {
			i = s.trackImage.Disabled
		}
	} else if s.hovering {
		if s.trackImage.Hover != nil {
			i = s.trackImage.Hover
		}
	}

	if i != nil {
		i.Draw(screen, s.widget.Rect.Dx(), s.widget.Rect.Dy(), s.widget.drawImageOptions)
	}
}

func (s *Slider) fireEvents() {
	if s.Current != s.lastCurrent {
		s.ChangedEvent.Fire(&SliderChangedEventArgs{
			Slider:  s,
			Current: s.Current,
		})
	}
}

func (s *Slider) updateHandleSize(handleLength float64) {
	l := int(math.Round(handleLength))
	if l < s.handleSize {
		l = s.handleSize
	}

	rect := s.widget.Rect

	var p img.Point
	if s.direction == DirectionHorizontal {
		p = img.Point{l, rect.Dy() - s.trackPadding.Top - s.trackPadding.Bottom}
	} else {
		p = img.Point{rect.Dx() - s.trackPadding.Left - s.trackPadding.Right, l}
	}

	s.handle.GetWidget().Rect.Max = s.handle.GetWidget().Rect.Min.Add(p)
}

func (s *Slider) updateHandleLocation(handleLength float64, trackLength float64) {
	internalTrackLength := int(math.Ceil(trackLength - handleLength))
	internalTrackStart := int(math.Floor(handleLength / 2))
	internalTrackEnd := internalTrackStart + internalTrackLength

	var i float64
	if s.dragging {
		x, y := input.CursorPosition()

		var dragOffset int
		if s.direction == DirectionHorizontal {
			dragOffset = x - s.handlePressedCursorX
		} else {
			dragOffset = y - s.handlePressedCursorY
		}
		var internalDragOffset float64
		if internalTrackLength > 0 {
			internalDragOffset = float64(dragOffset) / float64(internalTrackLength)
		} else {
			internalDragOffset = 0
		}

		i = s.handlePressedInternalCurrent + internalDragOffset
		if i < 0 {
			i = 0
		} else if i > 1 {
			i = 1
		}
		s.Current = s.internalToCurrent(i)
	} else {
		i = s.currentToInternal(s.Current)
	}

	off := int(math.Round(float64(internalTrackStart)*(1-i)+float64(internalTrackEnd)*i) - handleLength/2)

	rect := s.widget.Rect
	if s.direction == DirectionHorizontal {
		rect.Min = rect.Min.Add(img.Point{off + s.trackPadding.Left, s.trackPadding.Top})
	} else {
		rect.Min = rect.Min.Add(img.Point{s.trackPadding.Left, off + s.trackPadding.Top})
	}
	s.handle.GetWidget().Rect = rect
}

func (s *Slider) handleLengthAndTrackLength() (float64, float64) {
	var trackLength float64
	if s.direction == DirectionHorizontal {
		trackLength = float64(s.widget.Rect.Dx()) - float64(s.trackPadding.Left) - float64(s.trackPadding.Right)
	} else {
		trackLength = float64(s.widget.Rect.Dy()) - float64(s.trackPadding.Top) - float64(s.trackPadding.Bottom)
	}

	length := float64(s.Max - s.Min + 1)

	ps := s.pageSizeFunc()
	handleLength := float64(ps) / length * trackLength
	if handleLength < float64(s.handleSize) {
		handleLength = float64(s.handleSize)
	}
	if handleLength > trackLength {
		handleLength = trackLength
	}

	return handleLength, trackLength
}

func (s *Slider) currentToInternal(c int) float64 {
	if s.Max <= s.Min {
		return 0
	}

	return float64(c-s.Min) / float64(s.Max-s.Min)
}

func (s *Slider) internalToCurrent(i float64) int {
	return int(math.Round(float64(s.Min)*(1-i) + float64(s.Max)*i))
}

func (s *Slider) clampCurrentMinMax() {
	if s.Current < s.Min {
		s.Current = s.Min
	} else if s.Current > s.Max {
		s.Current = s.Max
	}
}

func (s *Slider) createWidget() {
	s.widget = NewWidget(append(s.widgetOpts, []WidgetOpt{
		WidgetOpts.CursorEnterHandler(func(args *WidgetCursorEnterEventArgs) {
			if !s.widget.Disabled {
				s.hovering = true
			}
		}),

		WidgetOpts.CursorExitHandler(func(args *WidgetCursorExitEventArgs) {
			s.hovering = false
		}),

		// TODO: keeping the mouse button pressed should move the handle repeatedly (in PageSize steps) until it stops under the cursor
		WidgetOpts.MouseButtonPressedHandler(func(args *WidgetMouseButtonPressedEventArgs) {
			if !s.widget.Disabled {
				x, y := input.CursorPosition()
				ps := s.pageSizeFunc()
				rect := s.handle.GetWidget().Rect
				if s.direction == DirectionHorizontal {
					if x < rect.Min.X {
						s.Current -= ps
					} else if x >= rect.Max.X {
						s.Current += ps
					}
				} else {
					if y < rect.Min.Y {
						s.Current -= ps
					} else if y >= rect.Max.Y {
						s.Current += ps
					}
				}

				s.clampCurrentMinMax()
			}
		}),
	}...)...)
	s.widgetOpts = nil

	s.handle = NewButton(append(s.handleOpts, []ButtonOpt{
		ButtonOpts.KeepPressedOnExit(),

		ButtonOpts.PressedHandler(func(args *ButtonPressedEventArgs) {
			s.dragging = true
			s.handlePressedCursorX, s.handlePressedCursorY = input.CursorPosition()
			s.handlePressedOffsetX = args.OffsetX
			s.handlePressedOffsetY = args.OffsetY
			s.handlePressedInternalCurrent = s.currentToInternal(s.Current)
		}),

		ButtonOpts.ReleasedHandler(func(args *ButtonReleasedEventArgs) {
			s.dragging = false
		}),
	}...)...)
}
