package widget

import (
	img "image"

	"github.com/hajimehoshi/ebiten/v2"
)

type TextToolTip struct {
	Label string

	containerOpts []ContainerOpt
	textOpts      []TextOpt
	padding       Insets

	init      *MultiOnce
	container *Container
	text      *Text
}

type TextToolTipOpt func(t *TextToolTip)

type TextToolTipOptions struct {
}

var TextToolTipOpts TextToolTipOptions

func NewTextToolTip(opts ...TextToolTipOpt) *TextToolTip {
	t := &TextToolTip{
		init: &MultiOnce{},
	}

	t.init.Append(t.createWidget)

	for _, o := range opts {
		o(t)
	}

	return t
}

func (o TextToolTipOptions) ContainerOpts(opts ...ContainerOpt) TextToolTipOpt {
	return func(t *TextToolTip) {
		t.containerOpts = append(t.containerOpts, opts...)
	}
}

func (o TextToolTipOptions) TextOpts(opts ...TextOpt) TextToolTipOpt {
	return func(t *TextToolTip) {
		t.textOpts = append(t.textOpts, opts...)
	}
}

func (o TextToolTipOptions) Padding(i Insets) TextToolTipOpt {
	return func(t *TextToolTip) {
		t.padding = i
	}
}

func (t *TextToolTip) GetWidget() *Widget {
	t.init.Do()
	return t.container.GetWidget()
}

func (t *TextToolTip) SetLocation(rect img.Rectangle) {
	t.init.Do()
	t.container.SetLocation(rect)
}

func (t *TextToolTip) PreferredSize() (int, int) {
	t.init.Do()

	t.text.Label = t.Label

	return t.container.PreferredSize()
}

func (t *TextToolTip) RequestRelayout() {
	t.init.Do()
	t.container.RequestRelayout()
}

func (t *TextToolTip) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	t.init.Do()

	t.text.Label = t.Label

	t.container.Render(screen, def, debugMode)
}

func (t *TextToolTip) createWidget() {
	t.container = NewContainer("text tooltip", append(t.containerOpts,
		ContainerOpts.Layout(NewAnchorLayout(
			AnchorLayoutOpts.Padding(t.padding),
		)),
	)...)

	t.text = NewText(t.textOpts...)
	t.text.Label = ""
	t.container.AddChild(t.text)
	t.textOpts = nil
}
