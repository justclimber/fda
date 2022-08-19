package widget

import (
	img "image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"

	"github.com/justclimber/fda/client/graphics/ui/event"
	"github.com/justclimber/fda/client/graphics/ui/image"
	"github.com/justclimber/fda/client/graphics/ui/input"
)

type List struct {
	EntrySelectedEvent *event.Event

	containerOpts            []ContainerOpt
	scrollContainerOpts      []ScrollContainerOpt
	sliderOpts               []SliderOpt
	entries                  []interface{}
	entryLabelFunc           ListEntryLabelFunc
	entryFace                font.Face
	entryUnselectedColor     *ButtonImage
	entrySelectedColor       *ButtonImage
	entryUnselectedTextColor *ButtonTextColor
	entryTextColor           *ButtonTextColor
	entryTextPadding         Insets
	controlWidgetSpacing     int
	hideHorizontalSlider     bool
	hideVerticalSlider       bool
	isMulti                  bool
	allowReselect            bool

	init            *MultiOnce
	container       *Container
	scrollContainer *ScrollContainer
	vSlider         *Slider
	hSlider         *Slider
	buttons         []*Button
	selectedEntry   interface{}
	selectedEntries []interface{}
}

type ListOpt func(l *List)

type ListEntryLabelFunc func(e interface{}) string

type ListEntryColor struct {
	Unselected                 color.Color
	Selected                   color.Color
	DisabledUnselected         color.Color
	DisabledSelected           color.Color
	SelectedBackground         color.Color
	DisabledSelectedBackground color.Color
}

type ListEntrySelectedEventArgs struct {
	List          *List
	Entry         interface{}
	PreviousEntry interface{}
}

type ListEntrySelectedHandlerFunc func(args *ListEntrySelectedEventArgs)

type ListOptions struct {
}

var ListOpts ListOptions

func NewList(opts ...ListOpt) *List {
	l := &List{
		EntrySelectedEvent: &event.Event{},

		init: &MultiOnce{},
	}

	l.init.Append(l.createWidget)

	for _, o := range opts {
		o(l)
	}

	return l
}

func (o ListOptions) ContainerOpts(opts ...ContainerOpt) ListOpt {
	return func(l *List) {
		l.containerOpts = append(l.containerOpts, opts...)
	}
}

func (o ListOptions) ScrollContainerOpts(opts ...ScrollContainerOpt) ListOpt {
	return func(l *List) {
		l.scrollContainerOpts = append(l.scrollContainerOpts, opts...)
	}
}

func (o ListOptions) SliderOpts(opts ...SliderOpt) ListOpt {
	return func(l *List) {
		l.sliderOpts = append(l.sliderOpts, opts...)
	}
}

func (o ListOptions) ControlWidgetSpacing(s int) ListOpt {
	return func(l *List) {
		l.controlWidgetSpacing = s
	}
}

func (o ListOptions) HideHorizontalSlider() ListOpt {
	return func(l *List) {
		l.hideHorizontalSlider = true
	}
}

func (o ListOptions) HideVerticalSlider() ListOpt {
	return func(l *List) {
		l.hideVerticalSlider = true
	}
}

func (o ListOptions) Entries(e []interface{}) ListOpt {
	return func(l *List) {
		l.entries = e
	}
}

func (o ListOptions) EntryLabelFunc(f ListEntryLabelFunc) ListOpt {
	return func(l *List) {
		l.entryLabelFunc = f
	}
}

func (o ListOptions) EntryFontFace(f font.Face) ListOpt {
	return func(l *List) {
		l.entryFace = f
	}
}

func (o ListOptions) EntryColor(c *ListEntryColor) ListOpt {
	return func(l *List) {
		l.entryUnselectedColor = &ButtonImage{
			Idle:     image.NewNineSliceColor(color.Transparent),
			Disabled: image.NewNineSliceColor(color.Transparent),
		}

		l.entrySelectedColor = &ButtonImage{
			Idle:     image.NewNineSliceColor(c.SelectedBackground),
			Disabled: image.NewNineSliceColor(c.DisabledSelectedBackground),
		}

		l.entryUnselectedTextColor = &ButtonTextColor{
			Idle:     c.Unselected,
			Disabled: c.DisabledUnselected,
		}

		l.entryTextColor = &ButtonTextColor{
			Idle:     c.Selected,
			Disabled: c.DisabledSelected,
		}
	}
}

func (o ListOptions) EntryTextPadding(i Insets) ListOpt {
	return func(l *List) {
		l.entryTextPadding = i
	}
}

func (o ListOptions) EntrySelectedHandler(f ListEntrySelectedHandlerFunc) ListOpt {
	return func(l *List) {
		l.EntrySelectedEvent.AddHandler(func(args interface{}) {
			f(args.(*ListEntrySelectedEventArgs))
		})
	}
}

func (o ListOptions) AllowReselect() ListOpt {
	return func(l *List) {
		l.allowReselect = true
	}
}

func (o ListOptions) IsMulti() ListOpt {
	return func(l *List) {
		l.isMulti = true
	}
}

func (l *List) GetWidget() *Widget {
	l.init.Do()
	return l.container.GetWidget()
}

func (l *List) PreferredSize() (int, int) {
	l.init.Do()
	return l.container.PreferredSize()
}

func (l *List) SetLocation(rect img.Rectangle) {
	l.init.Do()
	l.container.GetWidget().Rect = rect
}

func (l *List) RequestRelayout() {
	l.init.Do()
	l.container.RequestRelayout()
}

func (l *List) SetupInputLayer(def input.DeferredSetupInputLayerFunc) {
	l.init.Do()
	l.container.SetupInputLayer(def)
}

func (l *List) Render(screen *ebiten.Image, def DeferredRenderFunc, debugMode DebugMode) {
	l.init.Do()

	d := l.container.GetWidget().Disabled

	if l.vSlider != nil {
		l.vSlider.DrawTrackDisabled = d
	}
	if l.hSlider != nil {
		l.hSlider.DrawTrackDisabled = d
	}

	l.scrollContainer.GetWidget().Disabled = d

	l.container.Render(screen, def, debugMode)
}

func (l *List) createWidget() {
	var cols int
	if l.hideVerticalSlider {
		cols = 1
	} else {
		cols = 2
	}

	l.container = NewContainer(
		"list",
		append(l.containerOpts,
			ContainerOpts.Layout(NewGridLayout(
				GridLayoutOpts.Columns(cols),
				GridLayoutOpts.Stretch([]bool{true, false}, []bool{true, false}),
				GridLayoutOpts.Spacing(l.controlWidgetSpacing, l.controlWidgetSpacing))))...)
	l.containerOpts = nil

	content := NewContainer(
		"list content",
		ContainerOpts.Layout(NewRowLayout(
			RowLayoutOpts.Direction(DirectionVertical))),
		ContainerOpts.AutoDisableChildren())

	l.buttons = make([]*Button, 0, len(l.entries))
	for _, e := range l.entries {
		e := e
		but := NewButton(
			ButtonOpts.WidgetOpts(WidgetOpts.LayoutData(RowLayoutData{
				Stretch: true,
			})),
			ButtonOpts.Image(l.entryUnselectedColor),
			ButtonOpts.TextSimpleLeft(l.entryLabelFunc(e), l.entryFace, l.entryUnselectedTextColor, l.entryTextPadding),
			ButtonOpts.ClickedHandler(func(args *ButtonClickedEventArgs) {
				l.setSelectedEntry(e, true)
			}))

		l.buttons = append(l.buttons, but)

		content.AddChild(but)
	}

	l.scrollContainer = NewScrollContainer(append(l.scrollContainerOpts, []ScrollContainerOpt{
		ScrollContainerOpts.Content(content),
		ScrollContainerOpts.StretchContentWidth(),
	}...)...)
	l.scrollContainerOpts = nil
	l.container.AddChild(l.scrollContainer)

	if !l.hideVerticalSlider {
		pageSizeFunc := func() int {
			return int(math.Round(float64(l.scrollContainer.ContentRect().Dy()) / float64(content.GetWidget().Rect.Dy()) * 1000))
		}

		l.vSlider = NewSlider(append(l.sliderOpts, []SliderOpt{
			SliderOpts.Direction(DirectionVertical),
			SliderOpts.MinMax(0, 1000),
			SliderOpts.PageSizeFunc(pageSizeFunc),
			SliderOpts.ChangedHandler(func(args *SliderChangedEventArgs) {
				l.scrollContainer.ScrollTop = float64(args.Slider.Current) / 1000
			}),
		}...)...)
		l.container.AddChild(l.vSlider)

		l.scrollContainer.widget.ScrolledEvent.AddHandler(func(args interface{}) {
			a := args.(*WidgetScrolledEventArgs)
			p := pageSizeFunc() / 3
			if p < 1 {
				p = 1
			}
			l.vSlider.Current -= int(math.Round(a.Y * float64(p)))
		})
	}

	if !l.hideHorizontalSlider {
		l.hSlider = NewSlider(append(l.sliderOpts, []SliderOpt{
			SliderOpts.Direction(DirectionHorizontal),
			SliderOpts.MinMax(0, 1000),
			SliderOpts.PageSizeFunc(func() int {
				return int(math.Round(float64(l.scrollContainer.ContentRect().Dx()) / float64(content.GetWidget().Rect.Dx()) * 1000))
			}),
			SliderOpts.ChangedHandler(func(args *SliderChangedEventArgs) {
				l.scrollContainer.ScrollLeft = float64(args.Slider.Current) / 1000
			}),
		}...)...)
		l.container.AddChild(l.hSlider)
	}

	l.sliderOpts = nil
}

func (l *List) SetSelectedEntry(e interface{}) {
	l.setSelectedEntry(e, false)
}

func (l *List) setSelectedEntry(e interface{}, user bool) {
	if !l.isMulti && e == l.selectedEntry && (!user || !l.allowReselect) {
		return
	}
	l.init.Do()
	var prev interface{}
	if l.isMulti {
		if i, ok := inSelected(e, l.selectedEntries); ok {
			l.removeFromSelected(i)
		} else {
			l.selectedEntries = append(l.selectedEntries, e)
		}
		l.changeButtonsHighlight()
	} else {
		prev = l.selectedEntry
		l.selectedEntry = e

		l.changeButtonsHighlight()

	}
	l.EntrySelectedEvent.Fire(&ListEntrySelectedEventArgs{
		Entry:         e,
		PreviousEntry: prev,
	})
}

func (l *List) changeButtonsHighlight() {
	var selectedEntries []interface{}
	if l.isMulti {
		selectedEntries = l.selectedEntries
	} else {
		selectedEntries = []interface{}{l.selectedEntry}
	}

	for i, b := range l.buttons {
		if _, ok := inSelected(l.entries[i], selectedEntries); ok {
			b.Image = l.entrySelectedColor
			b.TextColor = l.entryTextColor
		} else {
			b.Image = l.entryUnselectedColor
			b.TextColor = l.entryUnselectedTextColor
		}
	}
}

func inSelected(e interface{}, l []interface{}) (int, bool) {
	for i, es := range l {
		if es == e {
			return i, true
		}
	}
	return 0, false
}

func (l *List) removeFromSelected(i int) {
	l.selectedEntries = append(l.selectedEntries[:i], l.selectedEntries[i+1:]...)
}

func (l *List) SelectedEntry() interface{} {
	l.init.Do()
	return l.selectedEntry
}

func (l *List) SelectedEntries() []interface{} {
	l.init.Do()
	return l.selectedEntries
}

func (l *List) SetScrollTop(t float64) {
	l.init.Do()
	if l.vSlider != nil {
		l.vSlider.Current = int(math.Round(t * 1000))
	}
	l.scrollContainer.ScrollTop = t
}

func (l *List) SetScrollLeft(left float64) {
	l.init.Do()
	if l.hSlider != nil {
		l.hSlider.Current = int(math.Round(left * 1000))
	}
	l.scrollContainer.ScrollLeft = left
}
