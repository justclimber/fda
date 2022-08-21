package widget

import (
	img "image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"github.com/justclimber/fda/client/graphics/ui/event"
	"github.com/justclimber/fda/client/graphics/ui/image"
	"github.com/justclimber/fda/client/graphics/ui/input"
)

type Button struct {
	Image             *ButtonImage
	KeepPressedOnExit bool
	GraphicImage      *ButtonImageImage
	TextColor         *ButtonTextColor

	PressedEvent  *event.Event
	ReleasedEvent *event.Event
	ClickedEvent  *event.Event

	widgetOpts               []Opt
	autoUpdateTextAndGraphic bool
	textPadding              Insets
	graphicPadding           Insets

	init      *MultiOnce
	widget    *Widget
	container *Container
	graphic   *Graphic
	text      *Text
	hovering  bool
	pressing  bool
}

type ButtonOpt func(b *Button)

type ButtonImage struct {
	Idle     *image.NineSlice
	Hover    *image.NineSlice
	Pressed  *image.NineSlice
	Disabled *image.NineSlice
}

type ButtonImageImage struct {
	Idle     *ebiten.Image
	Disabled *ebiten.Image
}

type ButtonTextColor struct {
	Idle     color.Color
	Disabled color.Color
}

type ButtonPressedEventArgs struct {
	Button  *Button
	OffsetX int
	OffsetY int
}

type ButtonReleasedEventArgs struct {
	Button  *Button
	Inside  bool
	OffsetX int
	OffsetY int
}

type ButtonClickedEventArgs struct {
	Button *Button
}

type ButtonPressedHandlerFunc func(args *ButtonPressedEventArgs)

type ButtonReleasedHandlerFunc func(args *ButtonReleasedEventArgs)

type ButtonClickedHandlerFunc func(args *ButtonClickedEventArgs)

type ButtonOptions struct {
}

var ButtonOpts ButtonOptions

func NewButton(opts ...ButtonOpt) *Button {
	b := &Button{
		PressedEvent:  &event.Event{},
		ReleasedEvent: &event.Event{},
		ClickedEvent:  &event.Event{},

		init: &MultiOnce{},
	}

	b.init.Append(b.createWidget)

	for _, o := range opts {
		o(b)
	}

	return b
}

func (o ButtonOptions) WidgetOpts(opts ...Opt) ButtonOpt {
	return func(b *Button) {
		b.widgetOpts = append(b.widgetOpts, opts...)
	}
}

func (o ButtonOptions) Image(i *ButtonImage) ButtonOpt {
	return func(b *Button) {
		b.Image = i
	}
}

func (o ButtonOptions) TextSimpleLeft(label string, face font.Face, color *ButtonTextColor, padding Insets) ButtonOpt {
	return func(b *Button) {
		b.init.Append(func() {
			b.container = NewContainer(
				"text simple left",
				ContainerOpts.Layout(NewAnchorLayout(AnchorLayoutOpts.Padding(padding))),
				ContainerOpts.AutoDisableChildren(),
			)

			b.text = NewText(
				TextOpts.WidgetOpts(Opts.LayoutData(AnchorLayoutData{
					HorizontalPosition: AnchorLayoutPositionStart,
					VerticalPosition:   AnchorLayoutPositionCenter,
				})),
				TextOpts.Text(label, face, color.Idle),
				TextOpts.Position(TextPositionStart, TextPositionCenter),
			)
			b.container.AddChild(b.text)

			b.autoUpdateTextAndGraphic = true
			b.TextColor = color
		})
	}
}

func (o ButtonOptions) Text(label string, face font.Face, color *ButtonTextColor) ButtonOpt {
	return func(b *Button) {
		b.init.Append(func() {
			b.container = NewContainer(
				"button text",
				ContainerOpts.Layout(NewAnchorLayout(AnchorLayoutOpts.Padding(b.textPadding))),
				ContainerOpts.AutoDisableChildren(),
			)

			b.text = NewText(
				TextOpts.WidgetOpts(Opts.LayoutData(AnchorLayoutData{
					HorizontalPosition: AnchorLayoutPositionCenter,
					VerticalPosition:   AnchorLayoutPositionCenter,
				})),
				TextOpts.Text(label, face, color.Idle),
				TextOpts.Position(TextPositionCenter, TextPositionCenter),
			)
			b.container.AddChild(b.text)

			b.autoUpdateTextAndGraphic = true
			b.TextColor = color
		})
	}
}

// TODO: add parameter for image position (start/end)
func (o ButtonOptions) TextAndImage(label string, face font.Face, image *ButtonImageImage, color *ButtonTextColor) ButtonOpt {
	return func(b *Button) {
		b.init.Append(func() {
			b.container = NewContainer(
				"button b",
				ContainerOpts.Layout(NewAnchorLayout(AnchorLayoutOpts.Padding(b.textPadding))),
				ContainerOpts.AutoDisableChildren(),
			)

			c := NewContainer(
				"button c",
				ContainerOpts.WidgetOpts(Opts.LayoutData(AnchorLayoutData{
					HorizontalPosition: AnchorLayoutPositionCenter,
					VerticalPosition:   AnchorLayoutPositionCenter,
				})),
				ContainerOpts.Layout(NewRowLayout(RowLayoutOpts.Spacing(10))),
				ContainerOpts.AutoDisableChildren(),
			)
			b.container.AddChild(c)

			b.text = NewText(
				TextOpts.WidgetOpts(Opts.LayoutData(RowLayoutData{
					Stretch: true,
				})),
				TextOpts.Text(label, face, color.Idle))
			c.AddChild(b.text)

			b.graphic = NewGraphic(
				GraphicOpts.WidgetOpts(Opts.LayoutData(RowLayoutData{
					Stretch: true,
				})),
				GraphicOpts.Image(image.Idle))
			c.AddChild(b.graphic)

			b.autoUpdateTextAndGraphic = true
			b.GraphicImage = image
			b.TextColor = color
		})
	}
}

func (o ButtonOptions) TextPadding(p Insets) ButtonOpt {
	return func(b *Button) {
		b.textPadding = p
	}
}

func (o ButtonOptions) Graphic(i *ebiten.Image) ButtonOpt {
	return o.withGraphic(GraphicOpts.Image(i))
}

func (o ButtonOptions) GraphicNineSlice(i *image.NineSlice) ButtonOpt {
	return o.withGraphic(GraphicOpts.ImageNineSlice(i))
}

func (o ButtonOptions) withGraphic(opt GraphicOpt) ButtonOpt {
	return func(b *Button) {
		b.init.Append(func() {
			b.container = NewContainer(
				"with graphic",
				ContainerOpts.Layout(NewAnchorLayout(AnchorLayoutOpts.Padding(b.graphicPadding))),
				ContainerOpts.AutoDisableChildren())

			b.graphic = NewGraphic(
				opt,
				GraphicOpts.WidgetOpts(Opts.LayoutData(AnchorLayoutData{
					HorizontalPosition: AnchorLayoutPositionCenter,
					VerticalPosition:   AnchorLayoutPositionCenter,
				})),
			)
			b.container.AddChild(b.graphic)

			b.autoUpdateTextAndGraphic = true
		})
	}
}

func (o ButtonOptions) GraphicPadding(i Insets) ButtonOpt {
	return func(b *Button) {
		b.graphicPadding = i
	}
}

func (o ButtonOptions) KeepPressedOnExit() ButtonOpt {
	return func(b *Button) {
		b.KeepPressedOnExit = true
	}
}

func (o ButtonOptions) PressedHandler(f ButtonPressedHandlerFunc) ButtonOpt {
	return func(b *Button) {
		b.PressedEvent.AddHandler(func(args interface{}) {
			f(args.(*ButtonPressedEventArgs))
		})
	}
}

func (o ButtonOptions) ReleasedHandler(f ButtonReleasedHandlerFunc) ButtonOpt {
	return func(b *Button) {
		b.ReleasedEvent.AddHandler(func(args interface{}) {
			f(args.(*ButtonReleasedEventArgs))
		})
	}
}

func (o ButtonOptions) ClickedHandler(f ButtonClickedHandlerFunc) ButtonOpt {
	return func(b *Button) {
		b.ClickedEvent.AddHandler(func(args interface{}) {
			f(args.(*ButtonClickedEventArgs))
		})
	}
}

func (b *Button) GetWidget() *Widget {
	b.init.Do()
	return b.widget
}

func (b *Button) PreferredSize() (int, int) {
	b.init.Do()

	w, h := 50, 50

	if b.container != nil && len(b.container.children) > 0 {
		w, h = b.container.PreferredSize()
	}

	iw, ih := b.Image.Idle.MinSize()
	if w < iw {
		w = iw
	}
	if h < ih {
		h = ih
	}

	return w, h
}

func (b *Button) SetLocation(rect img.Rectangle) {
	b.init.Do()
	b.widget.Rect = rect
}

func (b *Button) RequestRelayout() {
	b.init.Do()

	if b.container != nil {
		b.container.RequestRelayout()
	}
}

func (b *Button) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	b.init.Do()

	if b.container != nil {
		b.container.SetupInputLayer(def)
	}

	if b.pressing {
		def(func(def input.DeferredSetupInputLayerFunc) {
			b.widget.ElevateToNewInputLayer(&input.Layer{
				DebugLabel: "button pressed",
				EventTypes: input.LayerEventTypeAll,
				BlockLower: true,
				FullScreen: true,
			})
		})
	}
}

func (b *Button) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	b.init.Do()

	if b.container != nil {
		w := b.container.GetWidget()
		w.Rect = b.widget.Rect
		w.Disabled = b.widget.Disabled
	}

	b.widget.Render(screen, def, debugMode)

	b.draw(screen)

	if b.autoUpdateTextAndGraphic {
		if b.graphic != nil {
			if b.widget.Disabled {
				b.graphic.Image = b.GraphicImage.Disabled
			} else {
				b.graphic.Image = b.GraphicImage.Idle
			}
		}

		if b.text != nil {
			if b.widget.Disabled {
				b.text.Color = b.TextColor.Disabled
			} else {
				b.text.Color = b.TextColor.Idle
			}
		}
	}

	if b.container != nil {
		b.container.Render(screen, def, debugMode)
	}
}

func (b *Button) draw(screen *ebiten.Image) {
	i := b.Image.Idle
	switch {
	case b.widget.Disabled:
		if b.Image.Disabled != nil {
			i = b.Image.Disabled
		}
	case b.pressing && (b.hovering || b.KeepPressedOnExit):
		if b.Image.Pressed != nil {
			i = b.Image.Pressed
		}
	case b.hovering:
		if b.Image.Hover != nil {
			i = b.Image.Hover
		}
	}

	if i != nil {
		i.Draw(screen, b.widget.Rect.Dx(), b.widget.Rect.Dy(), func(opts *ebiten.DrawImageOptions) {
			b.widget.drawImageOptions(opts)
			b.drawImageOptions(opts)
		})
	}
}

func (b *Button) drawImageOptions(opts *ebiten.DrawImageOptions) {
	if b.widget.Disabled && b.Image.Disabled == nil {
		opts.ColorM.Scale(1, 1, 1, 0.35)
	}
}

func (b *Button) Text() *Text {
	b.init.Do()
	return b.text
}

func (b *Button) createWidget() {
	b.widget = NewWidget(append(b.widgetOpts, []Opt{
		Opts.CursorEnterHandler(func(args *CursorEnterEventArgs) {
			if !b.widget.Disabled {
				b.hovering = true
			}
		}),

		Opts.CursorExitHandler(func(args *CursorExitEventArgs) {
			b.hovering = false
		}),

		Opts.MouseButtonPressedHandler(func(args *MouseButtonPressedEventArgs) {
			if !b.widget.Disabled {
				b.pressing = true

				b.PressedEvent.Fire(&ButtonPressedEventArgs{
					Button:  b,
					OffsetX: args.OffsetX,
					OffsetY: args.OffsetY,
				})
			}
		}),

		Opts.MouseButtonReleasedHandler(func(args *MouseButtonReleasedEventArgs) {
			b.pressing = false

			if !b.widget.Disabled {
				b.ReleasedEvent.Fire(&ButtonReleasedEventArgs{
					Button:  b,
					Inside:  args.Inside,
					OffsetX: args.OffsetX,
					OffsetY: args.OffsetY,
				})

				if args.Inside {
					b.ClickedEvent.Fire(&ButtonClickedEventArgs{
						Button: b,
					})
				}
			}
		}),
	}...)...)
	b.widgetOpts = nil
}
