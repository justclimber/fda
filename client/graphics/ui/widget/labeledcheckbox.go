package widget

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/justclimber/fda/client/graphics/ui/input"
)

type LabeledCheckbox struct {
	checkboxOpts []CheckboxOpt
	labelOpts    []LabelOpt
	spacing      int

	init      *MultiOnce
	container *Container
	checkbox  *Checkbox
	label     *Label
}

type LabeledCheckboxOpt func(l *LabeledCheckbox)

type LabeledCheckboxOptions struct {
}

var LabeledCheckboxOpts LabeledCheckboxOptions

func NewLabeledCheckbox(opts ...LabeledCheckboxOpt) *LabeledCheckbox {
	l := &LabeledCheckbox{
		spacing: 8,

		init: &MultiOnce{},
	}

	l.init.Append(l.createWidget)

	for _, o := range opts {
		o(l)
	}

	return l
}

func (o LabeledCheckboxOptions) CheckboxOpts(opts ...CheckboxOpt) LabeledCheckboxOpt {
	return func(l *LabeledCheckbox) {
		l.checkboxOpts = append(l.checkboxOpts, opts...)
	}
}

func (o LabeledCheckboxOptions) LabelOpts(opts ...LabelOpt) LabeledCheckboxOpt {
	return func(l *LabeledCheckbox) {
		l.labelOpts = append(l.labelOpts, opts...)
	}
}

func (o LabeledCheckboxOptions) Spacing(s int) LabeledCheckboxOpt {
	return func(l *LabeledCheckbox) {
		l.spacing = s
	}
}

func (l *LabeledCheckbox) GetWidget() *Widget {
	l.init.Do()
	return l.container.GetWidget()
}

func (l *LabeledCheckbox) PreferredSize() (int, int) {
	l.init.Do()
	return l.container.PreferredSize()
}

func (l *LabeledCheckbox) RequestRelayout() {
	l.init.Do()
	l.container.RequestRelayout()
}

func (l *LabeledCheckbox) SetLocation(rect image.Rectangle) {
	l.init.Do()
	l.container.SetLocation(rect)
}

func (l *LabeledCheckbox) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	l.init.Do()
	l.checkbox.SetupInputLayer(def)
}

func (l *LabeledCheckbox) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	l.init.Do()
	l.container.Render(screen, def, debugMode)
}

func (l *LabeledCheckbox) Checkbox() *Checkbox {
	return l.checkbox
}

func (l *LabeledCheckbox) Label() *Label {
	return l.label
}

func (l *LabeledCheckbox) createWidget() {
	l.container = NewContainer(
		"labeled checkbox",
		ContainerOpts.Layout(NewRowLayout(
			RowLayoutOpts.Spacing(l.spacing))),
		ContainerOpts.AutoDisableChildren(),
	)

	l.checkbox = NewCheckbox(append(l.checkboxOpts, CheckboxOpts.ButtonOpts(ButtonOpts.WidgetOpts(WidgetOpts.LayoutData(RowLayoutData{
		Position: RowLayoutPositionCenter,
	}))))...)
	l.container.AddChild(l.checkbox)
	l.checkboxOpts = nil

	l.label = NewLabel(append(l.labelOpts, LabelOpts.TextOpts(TextOpts.WidgetOpts(
		WidgetOpts.LayoutData(RowLayoutData{
			Position: RowLayoutPositionCenter,
		}),

		WidgetOpts.MouseButtonReleasedHandler(func(args *WidgetMouseButtonReleasedEventArgs) {
			if !args.Widget.Disabled && args.Button == ebiten.MouseButtonLeft && args.Inside {
				l.checkbox.SetState(l.checkbox.state.Advance(l.checkbox.triState))
			}
		}),
	)))...)
	l.container.AddChild(l.label)
	l.labelOpts = nil
}
